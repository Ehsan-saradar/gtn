package types

import (
	"crypto/sha256"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(
	creator string,
	commitmentHash []byte,
	duration int64,
	entryFee sdk.Coin,
	reward sdk.Coin,
) *MsgCreateGame {
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

	if len(msg.CommitmentHash) != sha256.Size {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "commitment hash must be %d bytes long", sha256.Size)
	}

	if msg.Duration <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duration must be positive")
	}

	if err := msg.EntryFee.Validate(); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid entry fee (%s)", err)
	}

	if err := msg.Reward.Validate(); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid reward (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgRevealGame{}

func NewMsgRevealGame(
	creator string,
	gameId uint64,
	salt []byte,
	number uint64,
) *MsgRevealGame {
	return &MsgRevealGame{
		Creator: creator,
		GameId:  gameId,
		Salt:    salt,
		Number:  number,
	}
}

func (msg *MsgRevealGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Salt) != 32 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "salt must be 32 bytes long")
	}

	return nil
}
