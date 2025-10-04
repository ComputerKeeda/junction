package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/airchains-network/junction/x/evmbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) LockTokens(goCtx context.Context, msg *types.MsgLockTokens) (*types.MsgLockTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate creator address
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid creator address: %s", err))
	}

	// Validate amount
	if msg.Amount == "" {
		return nil, status.Error(codes.InvalidArgument, "amount is required")
	}

	// Parse amount as coins
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, types.ErrInvalidAmount.Wrap(err.Error())
	}

	if !amount.IsValid() || amount.IsZero() {
		return nil, types.ErrInvalidAmount.Wrap("amount must be positive")
	}

	// Validate EVM address format
	if err := k.validateEVMAddress(msg.ToAddress); err != nil {
		return nil, err
	}

	// Check if user has sufficient balance
	userBalance := k.bankKeeper.SpendableCoins(ctx, creatorAddr)
	if !userBalance.IsAllGTE(amount) {
		return nil, types.ErrInsufficientFunds.Wrapf("required: %s, available: %s", amount.String(), userBalance.String())
	}
	// Get module account
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		return nil, types.ErrModuleAccount.Wrap("module account not found")
	}

	// Transfer tokens from user to module account (this locks them)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, amount); err != nil {
		return nil, types.ErrFailedToLockTokens.Wrap(err.Error())
	}

	// Get current block info
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	blockTime := sdkCtx.BlockTime()

	amountUint64, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid amount: %s", err))
	}

	if err := k.SetAddressLockedAmount(ctx, msg.Creator, amountUint64); err != nil {
		return nil, types.ErrFailedToLockTokens.Wrap(err.Error())
	}
	if err := k.SetAddressMapping(ctx, msg.ToAddress, msg.Creator); err != nil {
		return nil, types.ErrFailedToMapAddress.Wrap(err.Error())
	}

	// Emit simple event for bridge relayer
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTokensLocked,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyToAddress, msg.ToAddress),
			sdk.NewAttribute(types.AttributeKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
			sdk.NewAttribute(types.AttributeKeyTimestamp, blockTime.Format(time.RFC3339)),
		),
	)

	return &types.MsgLockTokensResponse{}, nil
}

// validateEVMAddress performs basic EVM address validation
func (k msgServer) validateEVMAddress(address string) error {
	if address == "" {
		return types.ErrInvalidEVMAddr.Wrap("address cannot be empty")
	}

	// Basic EVM address validation (0x prefix + 40 hex chars)
	if len(address) != 42 {
		return types.ErrInvalidEVMAddr.Wrap("address must be 42 characters long")
	}

	if address[:2] != "0x" {
		return types.ErrInvalidEVMAddr.Wrap("address must start with 0x")
	}

	// Validate hex characters
	for i := 2; i < len(address); i++ {
		c := address[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return types.ErrInvalidEVMAddr.Wrap("address contains invalid hex characters")
		}
	}

	return nil
}
