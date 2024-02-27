package keeper

import (
	"bytes"
	"context"

	"gtn/x/gtn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var game = types.Game{
		Creator:         msg.Creator,
		CommitmentHash:  msg.CommitmentHash,
		StartedAtHeight: ctx.BlockHeight(),
		Duration:        msg.Duration,
		EntryFee:        msg.EntryFee,
		Reward:          msg.Reward,
	}

	id := k.AppendGame(
		ctx,
		game,
	)

	// Transfer the reward to the module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		sdk.MustAccAddressFromBech32(msg.Creator),
		types.ModuleName,
		sdk.NewCoins(msg.Reward),
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewCreateGameEvent(
			id,
			game.Creator,
			game.CommitmentHash,
			game.StartedAtHeight,
			game.Duration,
			game.EntryFee,
			game.Reward,
		),
	)

	return &types.MsgCreateGameResponse{
		Id: id,
	}, nil
}

func (k msgServer) RevealGame(goCtx context.Context, msg *types.MsgRevealGame) (*types.MsgRevealGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := k.GetGame(
		ctx,
		msg.GameId,
	)
	if !found {
		return nil, types.ErrGameNotFound
	}

	if game.Creator != msg.Creator {
		return nil, types.ErrUnauthorized
	}

	if game.Salt != nil {
		return nil, types.ErrGameAlreadyRevealed
	}

	if game.StartedAtHeight+game.Duration < ctx.BlockHeight() {
		return nil, types.ErrGameStillRunning
	}

	expDur := k.GetGameExpirationDuration(ctx)
	if game.StartedAtHeight+game.Duration+expDur > ctx.BlockHeight() {
		return nil, types.ErrGameExpired
	}

	commitment := types.CalculateCommitmentHash(msg.Salt, msg.Number)
	if bytes.Equal(commitment, game.CommitmentHash) {
		return nil, types.ErrInvalidCommitment
	}

	game.Salt = msg.Salt
	game.Number = msg.Number
	k.SetGame(
		ctx,
		game,
	)

	ctx.EventManager().EmitEvent(
		types.NewRevealGameEvent(
			game.Id,
			game.Creator,
			game.Salt,
			game.Number,
		),
	)

	return &types.MsgRevealGameResponse{}, nil
}
