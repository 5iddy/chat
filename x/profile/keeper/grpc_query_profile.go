package keeper

import (
	"context"

	"chat/x/profile/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProfileAll(c context.Context, req *types.QueryAllProfileRequest) (*types.QueryAllProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var profiles []types.Profile
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, types.KeyPrefix(types.ProfileKeyPrefix))

	pageRes, err := query.Paginate(profileStore, req.Pagination, func(key []byte, value []byte) error {
		var profile types.Profile
		if err := k.cdc.Unmarshal(value, &profile); err != nil {
			return err
		}

		profiles = append(profiles, profile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProfileResponse{Profile: profiles, Pagination: pageRes}, nil
}

func (k Keeper) Profile(c context.Context, req *types.QueryGetProfileRequest) (*types.QueryGetProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProfile(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProfileResponse{Profile: val}, nil
}
