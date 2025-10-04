package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/evmbridge module sentinel errors
var (
	ErrInvalidSigner           = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                  = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidAmount           = sdkerrors.Register(ModuleName, 1102, "invalid token amount")
	ErrInvalidEVMAddr          = sdkerrors.Register(ModuleName, 1103, "invalid EVM address")
	ErrInsufficientFunds       = sdkerrors.Register(ModuleName, 1104, "insufficient funds")
	ErrModuleAccount           = sdkerrors.Register(ModuleName, 1105, "module account error")
	ErrUnauthorized            = sdkerrors.Register(ModuleName, 1106, "unauthorized operation")
	ErrInvalidRecipient        = sdkerrors.Register(ModuleName, 1107, "invalid recipient address")
	ErrInsufficientUserBalance = sdkerrors.Register(ModuleName, 1108, "user has insufficient locked balance")
	ErrFailedToLockTokens      = sdkerrors.Register(ModuleName, 1109, "failed to lock tokens")
	ErrFailedToMapAddress      = sdkerrors.Register(ModuleName, 1110, "failed to map address")
	ErrInvalidAddress          = sdkerrors.Register(ModuleName, 1111, "invalid address")
	ErrInvalidEvmTxHash        = sdkerrors.Register(ModuleName, 1112, "invalid evm tx hash")
	ErrFailedToUnlockTokens    = sdkerrors.Register(ModuleName, 1113, "failed to unlock tokens")
)
