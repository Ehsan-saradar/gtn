package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, commitmentHash string, duration uint64, entryFee uint64, reward uint64) *MsgCreateGame {
	return &MsgCreateGame{
		Creator:        creator,
		CommitmentHash: commitmentHash,
		Duration:       duration,
		EntryFee:       entryFee,
		Reward:         reward,
	}
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGame{}

func NewMsgUpdateGame(creator string, id uint64, commitmentHash string, duration uint64, entryFee uint64, reward uint64) *MsgUpdateGame {
	return &MsgUpdateGame{
		Id:             id,
		Creator:        creator,
		CommitmentHash: commitmentHash,
		Duration:       duration,
		EntryFee:       entryFee,
		Reward:         reward,
	}
}

func (msg *MsgUpdateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteGame{}

func NewMsgDeleteGame(creator string, id uint64) *MsgDeleteGame {
	return &MsgDeleteGame{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
