package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTrackEngage{}

func NewMsgTrackEngage(creator string, trackKey string, schemaObject []byte) *MsgTrackEngage {
	return &MsgTrackEngage{
		Creator:      creator,
		TrackKey:     trackKey,
		SchemaObject: schemaObject,
	}
}

func (msg *MsgTrackEngage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
