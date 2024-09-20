package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetEspressoFinalizedState(goCtx context.Context, msg *types.MsgSetEspressoFinalizedState) (*types.MsgSetEspressoFinalizedStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetEspressoFinalizedStateResponse{}, nil
}
