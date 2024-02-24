package keeper

import (
	"context"
	"fmt"

	"gtn/x/gtn/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var game = types.Game{
		Creator:        msg.Creator,
		CommitmentHash: msg.CommitmentHash,
		Duration:       msg.Duration,
		EntryFee:       msg.EntryFee,
		Reward:         msg.Reward,
	}

	id := k.AppendGame(
		ctx,
		game,
	)

	return &types.MsgCreateGameResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateGame(goCtx context.Context, msg *types.MsgUpdateGame) (*types.MsgUpdateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var game = types.Game{
		Creator:        msg.Creator,
		Id:             msg.Id,
		CommitmentHash: msg.CommitmentHash,
		Duration:       msg.Duration,
		EntryFee:       msg.EntryFee,
		Reward:         msg.Reward,
	}

	// Checks that the element exists
	val, found := k.GetGame(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetGame(ctx, game)

	return &types.MsgUpdateGameResponse{}, nil
}

func (k msgServer) DeleteGame(goCtx context.Context, msg *types.MsgDeleteGame) (*types.MsgDeleteGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetGame(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveGame(ctx, msg.Id)

	return &types.MsgDeleteGameResponse{}, nil
}
