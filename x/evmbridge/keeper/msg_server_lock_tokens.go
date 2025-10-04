package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/evmbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LockTokens(goCtx context.Context, msg *types.MsgLockTokens) (*types.MsgLockTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLockTokensResponse{}, nil
}
