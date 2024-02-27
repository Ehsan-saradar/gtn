package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	"gtn/x/gtn/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

func (k Keeper) GetMaxPlayersPerGame(ctx context.Context) uint64 {
	return k.GetParams(ctx).MaxPlayersPerGame
}

func (k Keeper) SetMaxPlayersPerGame(ctx context.Context, maxPlayersPerGame uint64) {
	params := k.GetParams(ctx)
	params.MaxPlayersPerGame = maxPlayersPerGame
	k.SetParams(ctx, params)
}

func (k Keeper) GetMinDistanceToWin(ctx context.Context) uint64 {
	return k.GetParams(ctx).MinDistanceToWin
}

func (k Keeper) SetMinDistanceToWin(ctx context.Context, minDistanceToWin uint64) {
	params := k.GetParams(ctx)
	params.MinDistanceToWin = minDistanceToWin
	k.SetParams(ctx, params)
}

func (k Keeper) GetGameExpirationDuration(ctx context.Context) int64 {
	return k.GetParams(ctx).GameExpirationDuration
}

func (k Keeper) SetGameExpirationDuration(ctx context.Context, gameExpirationDuration int64) {
	params := k.GetParams(ctx)
	params.GameExpirationDuration = gameExpirationDuration
	k.SetParams(ctx, params)
}
