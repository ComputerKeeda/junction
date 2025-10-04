package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLockTokens{}

func NewMsgLockTokens(creator string, toAddress string, amount string) *MsgLockTokens {
	return &MsgLockTokens{
		Creator:   creator,
		ToAddress: toAddress,
		Amount:    amount,
	}
}

func (msg *MsgLockTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
