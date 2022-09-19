package keeper

import (
	"context"

	"chat/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddWebsiteToProfile(goCtx context.Context, msg *types.MsgAddWebsiteToProfile) (*types.MsgAddWebsiteToProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentProfile, isFound := k.GetProfile(ctx, msg.Name)
	if isFound && currentProfile.Creator == msg.Creator {
		profile := types.Profile{
			Name:    currentProfile.Name,
			Bio:     currentProfile.Bio,
			Website: msg.Website,
			Posts:   currentProfile.Posts,
			Creator: currentProfile.Creator,
		}
		k.SetProfile(ctx, profile)
		return &types.MsgAddWebsiteToProfileResponse{}, nil
	} else {
		err := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized Request")
		return nil, err
	}
}
