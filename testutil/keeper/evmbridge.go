package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address" // Import address codec
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/airchains-network/junction/x/evmbridge/keeper"
	"github.com/airchains-network/junction/x/evmbridge/types"
)

// Define necessary module account permissions for the test setup
var ModuleAccountPerms = []*authtypes.ModuleAccount{
	// Add module accounts used by your module or dependencies if any
	// For example: {Name: types.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
	{Name: types.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner, authtypes.Staking}},
}

func EvmbridgeKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Bech32 Prefix (Ensure this is set for your tests, e.g., in TestMain)
	// sdk.GetConfig().SetBech32Prefix("yourprefix")
	bech32Prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
	if bech32Prefix == "" {
		// Set a default if not configured for the test run
		bech32Prefix = "air"
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(bech32Prefix, bech32Prefix+"pub")
		config.SetBech32PrefixForValidator(bech32Prefix+"val", bech32Prefix+"valpub")
		config.SetBech32PrefixForConsensusNode(bech32Prefix+"cons", bech32Prefix+"conspub")
		config.Seal()
	}

	// Account Keeper (Needs authStoreKey, proper ModuleAccountPerms)
	// Create a map from the ModuleAccountPerms slice
	maccPerms := make(map[string][]string)
	for _, perm := range ModuleAccountPerms {
		maccPerms[perm.Name] = perm.Permissions
	}
	// If your module has its own module account, add it here too
	maccPerms[types.ModuleName] = nil // Example: No special permissions needed by default

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.NewBech32Codec(bech32Prefix),
		bech32Prefix,
		authority.String(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		accountKeeper,
		map[string]bool{},
		authority.String(),
		log.NewNopLogger())

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		accountKeeper,
		bankKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}
