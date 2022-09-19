package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddWebsiteToProfile = "add_website_to_profile"

var _ sdk.Msg = &MsgAddWebsiteToProfile{}

func NewMsgAddWebsiteToProfile(creator string, name string, website string) *MsgAddWebsiteToProfile {
	return &MsgAddWebsiteToProfile{
		Creator: creator,
		Name:    name,
		Website: website,
	}
}

func (msg *MsgAddWebsiteToProfile) Route() string {
	return RouterKey
}

func (msg *MsgAddWebsiteToProfile) Type() string {
	return TypeMsgAddWebsiteToProfile
}

func (msg *MsgAddWebsiteToProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddWebsiteToProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddWebsiteToProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
