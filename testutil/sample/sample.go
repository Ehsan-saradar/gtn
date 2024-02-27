package sample

import (
	"crypto/sha256"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// CommitmentHash returns a sample commitment hash
func CommitmentHash() []byte {
	hash := sha256.Sum256([]byte("commitment"))
	return hash[:]
}

// Coin returns a sample coin
func Coin() sdk.Coin {
	return sdk.NewCoin("stake", math.NewInt(100))
}
