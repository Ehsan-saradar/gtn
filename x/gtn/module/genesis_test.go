package gtn_test

import (
	"testing"

	keepertest "gtn/testutil/keeper"
	"gtn/testutil/nullify"
	gtn "gtn/x/gtn/module"
	"gtn/x/gtn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GameList: []types.Game{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GameCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GtnKeeper(t)
	gtn.InitGenesis(ctx, k, genesisState)
	got := gtn.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GameList, got.GameList)
	require.Equal(t, genesisState.GameCount, got.GameCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
