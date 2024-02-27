package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitGuess{}

func NewMsgCreateGuess(
	participant string,
	gameId uint64,
	number uint64,
) *MsgSubmitGuess {
	return &MsgSubmitGuess{
		Participant: participant,
		GameId:      gameId,
		Number:      number,
	}
}

func (msg *MsgSubmitGuess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
