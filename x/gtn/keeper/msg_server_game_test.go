package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"gtn/testutil/sample"
	"gtn/x/gtn/types"
)

func TestGameMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := sample.AccAddress()
	reward := sample.Coin()
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateGame(wctx, &types.MsgCreateGame{Creator: creator, Reward: reward})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestGameMsgServerReveal(t *testing.T) {
	creator := sample.AccAddress()

	tests := []struct {
		desc    string
		request *types.MsgRevealGame
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgRevealGame{
				Creator: creator,
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgRevealGame{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgRevealGame{Creator: creator, GameId: 10},
			err:     types.ErrGameNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateGame(wctx, &types.MsgCreateGame{
				Creator:        creator,
				CommitmentHash: sample.CommitmentHash(),
				Duration:       10,
				EntryFee:       sample.Coin(),
				Reward:         sample.Coin(),
			})
			require.NoError(t, err)

			_, err = srv.RevealGame(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
