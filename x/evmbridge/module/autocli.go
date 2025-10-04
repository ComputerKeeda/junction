package evmbridge

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/airchains-network/junction/api/junction/evmbridge"
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
					RpcMethod:      "LockTokens",
					Use:            "lock-tokens [to-address] [amount]",
					Short:          "Send a lock-tokens tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "toAddress"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "UnlockTokens",
					Use:            "unlock-tokens [to-address] [amount] [evm-tx-hash]",
					Short:          "Send a unlock-tokens tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "toAddress"}, {ProtoField: "amount"}, {ProtoField: "evmTxHash"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
