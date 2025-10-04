package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/airchains-network/junction/x/evmbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnlockTokens(goCtx context.Context, msg *types.MsgUnlockTokens) (*types.MsgUnlockTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allowedWorkers := k.GetParams(ctx).BridgeWorkers
	var isAllowedWorker bool
	for _, worker := range allowedWorkers {
		if worker == msg.Creator {
			isAllowedWorker = true
			break
		}
	}
	if !isAllowedWorker {
		return nil, types.ErrUnauthorized.Wrapf("creator %s is not authorized to unlock tokens", msg.Creator)
	}

	// get the cosmos address from the evm address
	cosmosAddress, err := k.GetCosmosAddressMapping(ctx, msg.FromEvmAddress)
	if err != nil {
		return nil, types.ErrInvalidAddress.Wrap(err.Error())
	}

	// check if the to address is a valid cosmos address
	_, err = sdk.AccAddressFromBech32(cosmosAddress)
	if err != nil {
		return nil, types.ErrInvalidAddress.Wrapf("invalid to address: %s", err)
	}

	// Validate evm tx hash
	if msg.EvmTxHash == "" {
		return nil, types.ErrInvalidEvmTxHash.Wrap("evm tx hash is required")
	}

	// we need to check if the passed amount is a valid number or not
	_, err = strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		return nil, types.ErrInvalidAmount.Wrapf("invalid amount: %s", err)
	}
	// now we need to check if the amount is less than or equal to the locked amount
	lockedAmount, err := k.GetAddressLockedAmount(ctx, cosmosAddress)
	if err != nil {
		return nil, types.ErrInvalidAddress.Wrap(err.Error())
	}
	// Parse string amount to sdk.Coins first
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, types.ErrInvalidAmount.Wrap(err.Error())
	}
	amountUint64, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		return nil, types.ErrInvalidAmount.Wrapf("invalid amount: %s", err)
	}
	if amountUint64 > lockedAmount {
		return nil, types.ErrInsufficientUserBalance.Wrapf("insufficient locked amount: %d", amountUint64)
	}

	// Get module account
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		return nil, types.ErrModuleAccount.Wrap("module account not found")
	}

	// Check if module has sufficient balance
	moduleBalance := k.bankKeeper.SpendableCoins(ctx, moduleAccount.GetAddress())
	if !moduleBalance.IsAllGTE(amount) {
		return nil, types.ErrInsufficientFunds.Wrapf("module insufficient funds: required %s, available %s", amount.String(), moduleBalance.String())
	}

	// transfer the amount from the module account to the to address
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(cosmosAddress), amount)
	if err != nil {
		return nil, types.ErrFailedToUnlockTokens.Wrap(err.Error())
	}

	// now we need to subtract the amount from the locked amount
	_, err = k.SubtractAddressLockedAmount(ctx, cosmosAddress, amountUint64)
	if err != nil {
		return nil, types.ErrFailedToUnlockTokens.Wrap(err.Error())
	}

	// Get current block info for event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	blockTime := sdkCtx.BlockTime()

	// Emit simple unlock event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTokensUnlocked,
			sdk.NewAttribute(types.AttributeKeyRecipient, cosmosAddress),
			sdk.NewAttribute(types.AttributeKeyFromEvmAddress, msg.FromEvmAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyEVMTxHash, msg.EvmTxHash),
			sdk.NewAttribute(types.AttributeKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
			sdk.NewAttribute(types.AttributeKeyTimestamp, blockTime.Format(time.RFC3339)),
		),
	)

	return &types.MsgUnlockTokensResponse{}, nil
}
