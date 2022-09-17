package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProfile = "create_profile"
	TypeMsgUpdateProfile = "update_profile"
	TypeMsgDeleteProfile = "delete_profile"
)

var _ sdk.Msg = &MsgCreateProfile{}

func NewMsgCreateProfile(
	creator string,
	name string,
	bio string,
	website string,
	posts []uint64,

) *MsgCreateProfile {
	return &MsgCreateProfile{
		Creator: creator,
		Name:    name,
		Bio:     bio,
		Website: website,
		Posts:   posts,
	}
}

func (msg *MsgCreateProfile) Route() string {
	return RouterKey
}

func (msg *MsgCreateProfile) Type() string {
	return TypeMsgCreateProfile
}

func (msg *MsgCreateProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProfile{}

func NewMsgUpdateProfile(
	creator string,
	name string,
	bio string,
	website string,
	posts []uint64,

) *MsgUpdateProfile {
	return &MsgUpdateProfile{
		Creator: creator,
		Name:    name,
		Bio:     bio,
		Website: website,
		Posts:   posts,
	}
}

func (msg *MsgUpdateProfile) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProfile) Type() string {
	return TypeMsgUpdateProfile
}

func (msg *MsgUpdateProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProfile{}

func NewMsgDeleteProfile(
	creator string,
	name string,

) *MsgDeleteProfile {
	return &MsgDeleteProfile{
		Creator: creator,
		Name:    name,
	}
}
func (msg *MsgDeleteProfile) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProfile) Type() string {
	return TypeMsgDeleteProfile
}

func (msg *MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
