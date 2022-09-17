package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateComment = "create_comment"
	TypeMsgUpdateComment = "update_comment"
	TypeMsgDeleteComment = "delete_comment"
)

var _ sdk.Msg = &MsgCreateComment{}

func NewMsgCreateComment(creator string, content string, replyTo string, createdAt int32, extention string, extensionType string) *MsgCreateComment {
	return &MsgCreateComment{
		Creator:       creator,
		Content:       content,
		ReplyTo:       replyTo,
		CreatedAt:     createdAt,
		Extention:     extention,
		ExtensionType: extensionType,
	}
}

func (msg *MsgCreateComment) Route() string {
	return RouterKey
}

func (msg *MsgCreateComment) Type() string {
	return TypeMsgCreateComment
}

func (msg *MsgCreateComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateComment{}

func NewMsgUpdateComment(creator string, id uint64, content string, replyTo string, createdAt int32, extention string, extensionType string) *MsgUpdateComment {
	return &MsgUpdateComment{
		Id:            id,
		Creator:       creator,
		Content:       content,
		ReplyTo:       replyTo,
		CreatedAt:     createdAt,
		Extention:     extention,
		ExtensionType: extensionType,
	}
}

func (msg *MsgUpdateComment) Route() string {
	return RouterKey
}

func (msg *MsgUpdateComment) Type() string {
	return TypeMsgUpdateComment
}

func (msg *MsgUpdateComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteComment{}

func NewMsgDeleteComment(creator string, id uint64) *MsgDeleteComment {
	return &MsgDeleteComment{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteComment) Route() string {
	return RouterKey
}

func (msg *MsgDeleteComment) Type() string {
	return TypeMsgDeleteComment
}

func (msg *MsgDeleteComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
