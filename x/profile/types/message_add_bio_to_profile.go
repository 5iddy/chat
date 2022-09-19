package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBioToProfile = "add_bio_to_profile"

var _ sdk.Msg = &MsgAddBioToProfile{}

func NewMsgAddBioToProfile(creator string, name string, bio string) *MsgAddBioToProfile {
	return &MsgAddBioToProfile{
		Creator: creator,
		Name:    name,
		Bio:     bio,
	}
}

func (msg *MsgAddBioToProfile) Route() string {
	return RouterKey
}

func (msg *MsgAddBioToProfile) Type() string {
	return TypeMsgAddBioToProfile
}

func (msg *MsgAddBioToProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBioToProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBioToProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
