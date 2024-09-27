package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TrackEngage(goCtx context.Context, msg *types.MsgTrackEngage) (*types.MsgTrackEngageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	trackKey := msg.TrackKey
	trackKeyBytes := []byte(trackKey)
	schemaObjectBytes := msg.SchemaObject

	// find the schemaDetails from the database
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TrackSchemaStoreKey))
	trackSchemaDetailsBytes := trackSchemaStore.Get(trackKeyBytes)

	if trackSchemaDetailsBytes == nil {
		return nil, status.Error(codes.NotFound, "track details not found")
	}

	var structDef types.StructDef
	if err := json.Unmarshal(trackSchemaDetailsBytes, &structDef); err != nil {
		return nil, err
	}

	schemaObjectString := string(schemaObjectBytes)

	_, err := DynamicUnmarshal(structDef, schemaObjectString)
	if err != nil {
		return nil, err
	}

	return &types.MsgTrackEngageResponse{
		Status: true,
	}, nil
}
