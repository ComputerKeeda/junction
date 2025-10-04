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

func (k Keeper) GetAddressLockedAmount(ctx sdk.Context, address string) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rollupRegistryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	amountBytes := rollupRegistryStore.Get([]byte(address))
	if amountBytes == nil {
		return 0, nil
	}
	amount := binary.BigEndian.Uint64(amountBytes)
	return amount, nil
}

func (k Keeper) AddAddressLockedAmount(ctx sdk.Context, address string, amountToAdd uint64) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rollupRegistryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Get current amount
	currentAmount, err := k.GetAddressLockedAmount(ctx, address)
	if err != nil {
		return 0, err
	}

	// Add the new amount
	newAmount := currentAmount + amountToAdd

	// Convert to bytes and store
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, newAmount)
	rollupRegistryStore.Set([]byte(address), amountBytes)

	return newAmount, nil
}

func (k Keeper) SubtractAddressLockedAmount(ctx sdk.Context, address string, amountToSubtract uint64) (uint64, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rollupRegistryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Get current amount
	currentAmount, err := k.GetAddressLockedAmount(ctx, address)
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
	rollupRegistryStore.Set([]byte(address), amountBytes)

	return newAmount, nil
}

func (k Keeper) SetAddressLockedAmount(ctx sdk.Context, address string, amount uint64) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rollupRegistryStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EVMLedgerKey))

	// Convert to bytes and store
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, amount)
	rollupRegistryStore.Set([]byte(address), amountBytes)

	return nil
}
