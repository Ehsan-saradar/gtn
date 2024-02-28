package keeper_test

import (
	"context"
	"testing"

	keepertest "gtn/testutil/keeper"
	"gtn/testutil/nullify"
	"gtn/x/gtn/keeper"
	"gtn/x/gtn/types"

	"github.com/stretchr/testify/require"
)

func createNGuess(keeper keeper.Keeper, ctx context.Context, n int) []types.Guess {
	items := make([]types.Guess, n)
	for i := range items {
		items[i].GameId = uint64(i)
		keeper.AppendGuess(ctx, items[i])
	}
	return items
}

func TestGuessGet(t *testing.T) {
	keeper, ctx := keepertest.GtnKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetGuess(ctx, item.GameId, "")
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestGuessRemove(t *testing.T) {
	keeper, ctx := keepertest.GtnKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGuess(ctx, item.GameId, "")
		_, found := keeper.GetGuess(ctx, item.GameId, "")
		require.False(t, found)
	}
}

func TestGuessGetAll(t *testing.T) {
	keeper, ctx := keepertest.GtnKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGuess(ctx)),
	)
}
