package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RetrieveTrackKey(goCtx context.Context, req *types.QueryRetrieveTrackKeyRequest) (*types.QueryRetrieveTrackKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	trackId := req.TrackId
	trackKeyFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackSchemaFinderStoreKey))
	// track id in bytes
	tidb := []byte(trackId)
	trackKeyBytes := trackKeyFinderStore.Get(tidb)

	if trackKeyBytes == nil {
		return &types.QueryRetrieveTrackKeyResponse{
			Track: nil,
		}, status.Error(codes.FailedPrecondition, "track key not found")
	}

	var extTrack types.ExtTrack
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackSchemaStoreKey))
	trackDetails := trackSchemaStore.Get(trackKeyBytes)
	if trackDetails == nil {
		return &types.QueryRetrieveTrackKeyResponse{
				Track: nil,
			},
			status.Error(codes.FailedPrecondition, "track detail not found")
	}

	k.cdc.MustUnmarshal(trackDetails, &extTrack)

	return &types.QueryRetrieveTrackKeyResponse{
		Track: &extTrack,
	}, nil
}
