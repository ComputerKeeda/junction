package junction

import (
	"math/rand"

	"github.com/airchains-network/junction/testutil/sample"
	junctionsimulation "github.com/airchains-network/junction/x/junction/simulation"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = junctionsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgInitStation = "op_weight_msg_init_station"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitStation int = 100

	opWeightMsgSubmitPod = "op_weight_msg_submit_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPod int = 100

	opWeightMsgVerifyPod = "op_weight_msg_verify_pod"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyPod int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	junctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&junctionGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgInitStation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitStation, &weightMsgInitStation, nil,
		func(_ *rand.Rand) {
			weightMsgInitStation = defaultWeightMsgInitStation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitStation,
		junctionsimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitPod int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitPod, &weightMsgSubmitPod, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPod = defaultWeightMsgSubmitPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPod,
		junctionsimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyPod int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVerifyPod, &weightMsgVerifyPod, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyPod = defaultWeightMsgVerifyPod
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyPod,
		junctionsimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitStation,
			defaultWeightMsgInitStation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgInitStation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitPod,
			defaultWeightMsgSubmitPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgSubmitPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyPod,
			defaultWeightMsgVerifyPod,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				junctionsimulation.SimulateMsgVerifyPod(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
