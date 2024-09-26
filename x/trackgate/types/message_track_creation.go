package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTrackCreation{}

func NewMsgTrackCreation(creator string, trackName string, trackId string, version string, schema []byte, status bool) *MsgTrackCreation {
	return &MsgTrackCreation{
		Creator:   creator,
		TrackName: trackName,
		TrackId:   trackId,
		Version:   version,
		Schema:    schema,
		Status:    status,
	}
}

func (msg *MsgTrackCreation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
