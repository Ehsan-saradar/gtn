package types

import (
	"testing"

	"gtn/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateGame
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateGame{
				Creator:        "invalid_address",
				CommitmentHash: sample.CommitmentHash(),
				Duration:       10,
				EntryFee:       sample.Coin(),
				Reward:         sample.Coin(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid commitment hash",
			msg: MsgCreateGame{
				Creator:        sample.AccAddress(),
				CommitmentHash: []byte{},
				Duration:       10,
				EntryFee:       sample.Coin(),
				Reward:         sample.Coin(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid duration",
			msg: MsgCreateGame{
				Creator:        sample.AccAddress(),
				CommitmentHash: sample.CommitmentHash(),
				Duration:       0,
				EntryFee:       sample.Coin(),
				Reward:         sample.Coin(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid entry fee",
			msg: MsgCreateGame{
				Creator:        sample.AccAddress(),
				CommitmentHash: sample.CommitmentHash(),
				Duration:       10,
				EntryFee:       sdk.Coin{},
				Reward:         sample.Coin(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid reward",
			msg: MsgCreateGame{
				Creator:        sample.AccAddress(),
				CommitmentHash: sample.CommitmentHash(),
				Duration:       10,
				EntryFee:       sample.Coin(),
				Reward:         sdk.Coin{},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid",
			msg: MsgCreateGame{
				Creator:        sample.AccAddress(),
				CommitmentHash: sample.CommitmentHash(),
				Duration:       10,
				EntryFee:       sample.Coin(),
				Reward:         sample.Coin(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
