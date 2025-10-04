package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/evmbridge module sentinel errors
var (
	ErrInvalidSigner           = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                  = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidAmount           = sdkerrors.Register(ModuleName, 1101, "invalid token amount")
	ErrInvalidEVMAddr          = sdkerrors.Register(ModuleName, 1102, "invalid EVM address")
	ErrInsufficientFunds       = sdkerrors.Register(ModuleName, 1103, "insufficient funds")
	ErrModuleAccount           = sdkerrors.Register(ModuleName, 1104, "module account error")
	ErrUnauthorized            = sdkerrors.Register(ModuleName, 1105, "unauthorized operation")
	ErrInvalidRecipient        = sdkerrors.Register(ModuleName, 1106, "invalid recipient address")
	ErrInsufficientUserBalance = sdkerrors.Register(ModuleName, 1107, "user has insufficient locked balance")
	ErrFailedToLockTokens      = sdkerrors.Register(ModuleName, 1108, "failed to lock tokens")
)
