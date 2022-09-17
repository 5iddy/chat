package keeper

import (
	"context"
	"fmt"

	"chat/x/comments/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var comment = types.Comment{
		Creator:       msg.Creator,
		Content:       msg.Content,
		ReplyTo:       msg.ReplyTo,
		CreatedAt:     msg.CreatedAt,
		Extention:     msg.Extention,
		ExtensionType: msg.ExtensionType,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	return &types.MsgCreateCommentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateComment(goCtx context.Context, msg *types.MsgUpdateComment) (*types.MsgUpdateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var comment = types.Comment{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Content:       msg.Content,
		ReplyTo:       msg.ReplyTo,
		CreatedAt:     msg.CreatedAt,
		Extention:     msg.Extention,
		ExtensionType: msg.ExtensionType,
	}

	// Checks that the element exists
	val, found := k.GetComment(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetComment(ctx, comment)

	return &types.MsgUpdateCommentResponse{}, nil
}

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetComment(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveComment(ctx, msg.Id)

	return &types.MsgDeleteCommentResponse{}, nil
}
