package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"gtn/testutil/sample"
	"gtn/x/gtn/types"
)

func TestGuessMsgServerSubmit(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	// Create a game
	creator := sample.AccAddress()
	reward := sample.Coin()
	resp, err := srv.CreateGame(wctx, &types.MsgCreateGame{Creator: creator, Reward: reward})
	require.NoError(t, err)

	participant := sample.AccAddress()
	for i := 0; i < 5; i++ {
		_, err := srv.SubmitGuess(wctx, &types.MsgSubmitGuess{Participant: participant, GameId: resp.Id, Number: uint64(i)})
		require.NoError(t, err)
	}
}
