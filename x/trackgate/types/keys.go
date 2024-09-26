package types

const (
	// ModuleName defines the module name
	ModuleName = "trackgate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trackgate"

	// TrackSchemaFinderStoreKey is actually key:track-id and value:track-key
	TrackSchemaFinderStoreKey = "track_schema_finder"
	// TrackSchemaStoreKey is actually key:track-key and value:schema_bytes
	TrackSchemaStoreKey = "track_schema"
)

var (
	ParamsKey = []byte("p_trackgate")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
