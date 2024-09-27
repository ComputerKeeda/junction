package trackgate

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/trackgate"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "RetrieveTrackKey",
					Use:            "retrieve-track-key [track-id]",
					Short:          "Query retrieve-track-key",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trackId"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "TrackCreation",
					Use:            "track-creation [track-name] [track-id] [version] [schema] [status]",
					Short:          "Send a track-creation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trackName"}, {ProtoField: "trackId"}, {ProtoField: "version"}, {ProtoField: "schema"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "TrackEngage",
					Use:            "track-engage [track-key] [schema-object]",
					Short:          "Send a track-engage tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trackKey"}, {ProtoField: "schemaObject"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
