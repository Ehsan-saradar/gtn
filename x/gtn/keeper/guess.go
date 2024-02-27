package keeper

import (
	"context"
	"encoding/binary"

	"gtn/x/gtn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// AppendGuess appends a guess in the store with a new id and update the count
func (k Keeper) AppendGuess(
	ctx context.Context,
	guess types.Guess,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	appendedValue := k.cdc.MustMarshal(&guess)
	store.Set(GetGuessIDBytes(guess.GameId, guess.Participant), appendedValue)
}

// SetGuess set a specific guess in the store
func (k Keeper) SetGuess(ctx context.Context, guess types.Guess) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	b := k.cdc.MustMarshal(&guess)
	store.Set(GetGuessIDBytes(guess.GameId, guess.Participant), b)
}

// GetGuess returns a guess from its id
func (k Keeper) GetGuess(ctx context.Context, gameId uint64, participant string) (val types.Guess, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	b := store.Get(GetGuessIDBytes(gameId, participant))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetGameGuesses returns all guess of a game
func (k Keeper) GetGameGuesses(ctx context.Context, gameId uint64) (list []types.Guess) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	iterator := storetypes.KVStorePrefixIterator(store, GetGuessIDBytes(gameId, ""))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Guess
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveGuess removes a guess from the store
func (k Keeper) RemoveGuess(ctx context.Context, gameId uint64, participant string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	store.Delete(GetGuessIDBytes(gameId, participant))
}

// RemoveGameGuesses removes all guess of a game from the store
func (k Keeper) RemoveGameGuesses(ctx context.Context, gameId uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	iterator := storetypes.KVStorePrefixIterator(store, GetGuessIDBytes(gameId, ""))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

// GetAllGuess returns all guess
func (k Keeper) GetAllGuess(ctx context.Context) (list []types.Guess) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Guess
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGuessIDBytes returns the byte representation of the ID
func GetGuessIDBytes(gameId uint64, participant string) []byte {
	bz := types.KeyPrefix(types.GuessKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, gameId)
	bz = append(bz, []byte("/")...)
	bz = append(bz, []byte(participant)...)
	return bz
}
