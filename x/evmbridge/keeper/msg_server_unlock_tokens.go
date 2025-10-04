package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/evmbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnlockTokens(goCtx context.Context, msg *types.MsgUnlockTokens) (*types.MsgUnlockTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnlockTokensResponse{}, nil
}
