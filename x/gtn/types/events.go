package types

import (
	"encoding/hex"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CreateGameEventType  = "create_game"
	RevealGameEventType  = "reveal_game"
	SumbitGuessEventType = "submit_guess"
	GameExpiredEventType = "game_expired"
)

func NewCreateGameEvent(
	id uint64,
	creator string,
	commitmentHash []byte,
	startedAtHeight int64,
	duration int64,
	entryFee sdk.Coin,
	reward sdk.Coin,
) sdk.Event {
	return sdk.NewEvent(
		CreateGameEventType,
		sdk.NewAttribute("id", fmt.Sprint(id)),
		sdk.NewAttribute("creator", creator),
		sdk.NewAttribute("commitment_hash", hex.EncodeToString(commitmentHash)),
		sdk.NewAttribute("started_at_height", fmt.Sprint(startedAtHeight)),
		sdk.NewAttribute("duration", fmt.Sprint(duration)),
		sdk.NewAttribute("entry_fee", entryFee.String()),
		sdk.NewAttribute("reward", reward.String()),
	)
}

func NewRevealGameEvent(
	id uint64,
	creator string,
	salt []byte,
	number uint64,
) sdk.Event {
	return sdk.NewEvent(
		RevealGameEventType,
		sdk.NewAttribute("id", fmt.Sprint(id)),
		sdk.NewAttribute("creator", creator),
		sdk.NewAttribute("salt", hex.EncodeToString(salt)),
		sdk.NewAttribute("number", fmt.Sprint(number)),
	)
}

func NewSubmitGuessEvent(
	participant string,
	gameId uint64,
	number uint64,
) sdk.Event {
	return sdk.NewEvent(
		SumbitGuessEventType,
		sdk.NewAttribute("participant", participant),
		sdk.NewAttribute("game_id", fmt.Sprint(gameId)),
		sdk.NewAttribute("number", fmt.Sprint(number)),
	)
}

func NewGameExpiredEvent(
	id uint64,
) sdk.Event {
	return sdk.NewEvent(
		GameExpiredEventType,
		sdk.NewAttribute("id", fmt.Sprint(id)),
	)
}
