package keeper

import (
	"context"

	"chat/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddBioToProfile(goCtx context.Context, msg *types.MsgAddBioToProfile) (*types.MsgAddBioToProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentProfile, isFound := k.GetProfile(ctx, msg.Name)

	if isFound && currentProfile.Creator == msg.Creator {
		profile := types.Profile{
			Name:    currentProfile.Name,
			Bio:     msg.Bio,
			Website: currentProfile.Website,
			Posts:   currentProfile.Posts,
			Creator: currentProfile.Creator,
		}

		k.SetProfile(ctx, profile)
		return &types.MsgAddBioToProfileResponse{}, nil
	} else {
		err := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized request")
		return nil, err
	}

}
