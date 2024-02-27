package types

import (
	"crypto/sha256"
	"encoding/binary"
)

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

const (
	GuessKey = "Guess/value/"
)

var (
	CreateGameGas = uint64(200000)
)

func CalculateCommitmentHash(salt []byte, number uint64) []byte {
	hash := sha256.New()
	hash.Write(salt)
	binary.Write(hash, binary.BigEndian, number)
	return hash.Sum(nil)
}
