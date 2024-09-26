package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TrackCreation(goCtx context.Context, msg *types.MsgTrackCreation) (*types.MsgTrackCreationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	trackId := msg.TrackId
	trackName := msg.TrackName
	trackCreator := msg.Creator
	version := msg.Version
	schema := msg.Schema

	creators := make([]string, 0)
	// Add trackCreator to the creators slice
	creators = append(creators, trackCreator)

	validateTrackIDRes := ValidateTrackID(trackId)
	if validateTrackIDRes {
		return &types.MsgTrackCreationResponse{
			TrackKey: "",
			Status:   false,
		}, status.Error(codes.FailedPrecondition, "invalid track key. Refer to the documentation for valid formats.")
	}

	//Store => key will be new trackID and value will be trackKey
	trackKeyFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackSchemaFinderStoreKey))
	// track id in bytes
	tidb := []byte(trackId)
	trackKeyBytes := trackKeyFinderStore.Get(tidb)

	if trackKeyBytes == nil {
		return &types.MsgTrackCreationResponse{
			TrackKey: "",
			Status:   false,
		}, status.Error(codes.FailedPrecondition, "track id already exists.")
	}

	// Generate UUID
	trackKey := uuid.New().String()

	newExeTracks := types.ExtTrack{
		TrackName: trackName,
		TrackId:   trackId,
		TrackKey:  trackKey,
		Version:   version,
		Schema:    schema,
		Creators:  creators,
	}

	storingData := k.cdc.MustMarshal(&newExeTracks)
	//Store => key:track-key and value:schema_bytes
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackSchemaStoreKey))
	trackSchemaStore.Set(trackKeyBytes, storingData)

	return &types.MsgTrackCreationResponse{
		TrackKey: trackKey,
		Status:   true,
	}, nil
}
