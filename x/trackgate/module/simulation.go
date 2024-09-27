package trackgate

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/airchains-network/junction/testutil/sample"
	trackgatesimulation "github.com/airchains-network/junction/x/trackgate/simulation"
	"github.com/airchains-network/junction/x/trackgate/types"
)

// avoid unused import issue
var (
	_ = trackgatesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgTrackCreation = "op_weight_msg_track_creation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTrackCreation int = 100

	opWeightMsgTrackEngage = "op_weight_msg_track_engage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTrackEngage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	trackgateGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&trackgateGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgTrackCreation int
	simState.AppParams.GetOrGenerate(opWeightMsgTrackCreation, &weightMsgTrackCreation, nil,
		func(_ *rand.Rand) {
			weightMsgTrackCreation = defaultWeightMsgTrackCreation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTrackCreation,
		trackgatesimulation.SimulateMsgTrackCreation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTrackEngage int
	simState.AppParams.GetOrGenerate(opWeightMsgTrackEngage, &weightMsgTrackEngage, nil,
		func(_ *rand.Rand) {
			weightMsgTrackEngage = defaultWeightMsgTrackEngage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTrackEngage,
		trackgatesimulation.SimulateMsgTrackEngage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgTrackCreation,
			defaultWeightMsgTrackCreation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgTrackCreation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTrackEngage,
			defaultWeightMsgTrackEngage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				trackgatesimulation.SimulateMsgTrackEngage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
