package types

const (
	// ModuleName defines the module name
	ModuleName = "gtn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gtn"
)

var (
	ParamsKey = []byte("p_gtn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GameKey      = "Game/value/"
	GameCountKey = "Game/count/"
)
