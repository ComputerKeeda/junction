package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/airchains-network/junction/x/evmbridge/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAddressLockedAmount(ctx sdk.Context, cosmos_address string) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	amountBytes := ledgerStore.Get([]byte(cosmos_address))
	if amountBytes == nil {
		return 0, nil
	}
	amount := binary.BigEndian.Uint64(amountBytes)
	return amount, nil
}

func (k Keeper) AddAddressLockedAmount(ctx sdk.Context, cosmos_address string, amountToAdd uint64) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Get current amount
	currentAmount, err := k.GetAddressLockedAmount(ctx, cosmos_address)
	if err != nil {
		return 0, err
	}

	// Add the new amount
	newAmount := currentAmount + amountToAdd

	// Convert to bytes and store
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, newAmount)
	ledgerStore.Set([]byte(cosmos_address), amountBytes)

	return newAmount, nil
}

func (k Keeper) SubtractAddressLockedAmount(ctx sdk.Context, cosmos_address string, amountToSubtract uint64) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Get current amount
	currentAmount, err := k.GetAddressLockedAmount(ctx, cosmos_address)
	if err != nil {
		return 0, err
	}

	// Check for underflow
	if currentAmount < amountToSubtract {
		return 0, status.Error(codes.InvalidArgument, fmt.Sprintf("insufficient locked amount: have %d, trying to subtract %d", currentAmount, amountToSubtract))
	}

	// Subtract the amount
	newAmount := currentAmount - amountToSubtract

	// Convert to bytes and store
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, newAmount)
	ledgerStore.Set([]byte(cosmos_address), amountBytes)

	return newAmount, nil
}

func (k Keeper) SetAddressLockedAmount(ctx sdk.Context, cosmos_address string, amount uint64) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Convert to bytes and store
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, amount)
	ledgerStore.Set([]byte(cosmos_address), amountBytes)

	return nil
}

// Store address mapping (evm address -> cosmos address)
func (k Keeper) SetAddressMapping(ctx sdk.Context, evm_address string, cosmos_address string) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerMappingStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMMappingKey))

	ledgerMappingStore.Set([]byte(evm_address), []byte(cosmos_address))
	return nil
}

func (k Keeper) GetCosmosAddressMapping(ctx sdk.Context, evm_address string) (string, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ledgerMappingStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMMappingKey))

	cosmos_address := ledgerMappingStore.Get([]byte(evm_address))
	if cosmos_address == nil {
		return "", status.Error(codes.NotFound, "address mapping not found")
	}
	return string(cosmos_address), nil
}
