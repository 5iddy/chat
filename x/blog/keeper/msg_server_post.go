package keeper

import (
	"context"
	"fmt"

	"chat/x/blog/types"
	profiletypes "chat/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	currentProfile, isFound := k.profileKeeper.GetProfile(ctx, msg.Name)
	var id uint64

	if isFound && currentProfile.Creator == msg.Creator {
		post.Name = currentProfile.Name
		id = k.AppendPost(
			ctx,
			post,
		)
		currentProfile.Posts = append(currentProfile.Posts, id)
		k.profileKeeper.SetProfile(ctx, currentProfile)
		return &types.MsgCreatePostResponse{
			Id: id,
		}, nil
	}
	if isFound && currentProfile.Creator != msg.Creator {
		err := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized Request")
		return nil, err
	}
	if !isFound {
		post.Name = msg.Name
		id = k.AppendPost(
			ctx,
			post,
		)
		profile := profiletypes.Profile{
			Name:    msg.Name,
			Creator: msg.Creator,
			Posts:   []uint64{id},
		}
		k.profileKeeper.SetProfile(ctx, profile)
		return &types.MsgCreatePostResponse{
			Id: id,
		}, nil
	}

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Checks that the element exists
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPost(ctx, post)

	return &types.MsgUpdatePostResponse{}, nil
}

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePost(ctx, msg.Id)

	return &types.MsgDeletePostResponse{}, nil
}
