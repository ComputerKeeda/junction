package keeper

import (
	"github.com/airchains-network/junction/x/evmbridge/types"
)

var _ types.QueryServer = Keeper{}
