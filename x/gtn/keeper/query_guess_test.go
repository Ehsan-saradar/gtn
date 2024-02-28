package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gtn/testutil/keeper"
	"gtn/testutil/nullify"
	"gtn/x/gtn/types"
)

func TestGamesGuessesQuery(t *testing.T) {
	keeper, ctx := keepertest.GtnKeeper(t)
	msgs := createNGuess(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGameGuessesRequest
		response *types.QueryGetGameGuessesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetGameGuessesRequest{GameId: msgs[0].GameId},
			response: &types.QueryGetGameGuessesResponse{Guesses: []types.Guess{msgs[0]}},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetGameGuessesRequest{GameId: msgs[1].GameId},
			response: &types.QueryGetGameGuessesResponse{Guesses: []types.Guess{msgs[1]}},
		},
		{
			desc:     "NoGuessFound",
			request:  &types.QueryGetGameGuessesRequest{GameId: uint64(len(msgs) + 1)},
			response: &types.QueryGetGameGuessesResponse{Guesses: []types.Guess{}},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GameGuesses(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestGuessQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GtnKeeper(t)
	msgs := createNGuess(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGuessRequest {
		return &types.QueryAllGuessRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GuessAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Guesses), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Guesses),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GuessAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Guesses), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Guesses),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GuessAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Guesses),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GuessAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
