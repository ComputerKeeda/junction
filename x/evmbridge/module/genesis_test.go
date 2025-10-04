package evmbridge_test

import (
	"testing"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	evmbridge "github.com/airchains-network/junction/x/evmbridge/module"
	"github.com/airchains-network/junction/x/evmbridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvmbridgeKeeper(t)
	evmbridge.InitGenesis(ctx, k, genesisState)
	got := evmbridge.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
