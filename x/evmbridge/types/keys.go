package types

const (
	// ModuleName defines the module name
	ModuleName = "evmbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_evmbridge"

	EVMLedgerKey = "evm_ledger"

	EVMMappingKey = "evm_mapping"
)

var (
	ParamsKey = []byte("p_evmbridge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
