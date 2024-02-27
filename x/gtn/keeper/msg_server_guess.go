package keeper

import (
	"context"

	"gtn/x/gtn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitGuess(goCtx context.Context, msg *types.MsgSubmitGuess) (*types.MsgSubmitGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, types.ErrGameNotFound
	}

	if game.StartedAtHeight+game.Duration > ctx.BlockHeight() {
		return nil, types.ErrGameExpired
	}

	if game.Number != 0 {
		return nil, types.ErrGameAlreadyRevealed
	}

	maxPlayers := k.GetMaxPlayersPerGame(ctx)
	if game.ParticipantsCount >= maxPlayers {
		return nil, types.ErrGameFull
	}

	// Transfer the entry fee to the module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		sdk.MustAccAddressFromBech32(msg.Participant),
		types.ModuleName,
		sdk.NewCoins(game.EntryFee),
	); err != nil {
		return nil, err
	}

	game.ParticipantsCount++
	k.SetGame(ctx, game)

	var guess = types.Guess{
		Participant: msg.Participant,
		GameId:      msg.GameId,
		Number:      msg.Number,
	}
	k.AppendGuess(
		ctx,
		guess,
	)

	ctx.EventManager().EmitEvent(
		types.NewSubmitGuessEvent(
			guess.Participant,
			guess.GameId,
			guess.Number,
		),
	)

	return &types.MsgSubmitGuessResponse{}, nil
}
