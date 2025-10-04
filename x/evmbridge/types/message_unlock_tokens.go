package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnlockTokens{}

func NewMsgUnlockTokens(creator string, toAddress string, amount string, evmTxHash string) *MsgUnlockTokens {
	return &MsgUnlockTokens{
		Creator:   creator,
		ToAddress: toAddress,
		Amount:    amount,
		EvmTxHash: evmTxHash,
	}
}

func (msg *MsgUnlockTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
