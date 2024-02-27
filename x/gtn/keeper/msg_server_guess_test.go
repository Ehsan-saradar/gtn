package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"gtn/x/gtn/types"
)

func TestGuessMsgServerSubmit(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	participant := "A"
	for i := 0; i < 5; i++ {
		_, err := srv.SubmitGuess(wctx, &types.MsgSubmitGuess{Participant: participant})
		require.NoError(t, err)
	}
}
