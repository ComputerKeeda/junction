package keeper_test

import (
	"testing"

	testkeeper "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JunctionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
