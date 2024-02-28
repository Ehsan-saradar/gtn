package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/gtn module sentinel errors
var (
	ErrInvalidSigner       = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrGameNotFound        = sdkerrors.Register(ModuleName, 1102, "game not found")
	ErrInvalidCommitment   = sdkerrors.Register(ModuleName, 1103, "invalid commitment")
	ErrGameAlreadyRevealed = sdkerrors.Register(ModuleName, 1104, "game already revealed")
	ErrGameExpired         = sdkerrors.Register(ModuleName, 1105, "game expired")
	ErrGameFull            = sdkerrors.Register(ModuleName, 1106, "game full")
	ErrGameStillRunning    = sdkerrors.Register(ModuleName, 1107, "game still running")
)
