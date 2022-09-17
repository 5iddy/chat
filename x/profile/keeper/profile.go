package keeper

import (
	"chat/x/profile/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetProfile set a specific profile in the store from its index
func (k Keeper) SetProfile(ctx sdk.Context, profile types.Profile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKeyPrefix))
	b := k.cdc.MustMarshal(&profile)
	store.Set(types.ProfileKey(
		profile.Name,
	), b)
}

// GetProfile returns a profile from its index
func (k Keeper) GetProfile(
	ctx sdk.Context,
	name string,

) (val types.Profile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKeyPrefix))

	b := store.Get(types.ProfileKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProfile removes a profile from the store
func (k Keeper) RemoveProfile(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKeyPrefix))
	store.Delete(types.ProfileKey(
		name,
	))
}

// GetAllProfile returns all profile
func (k Keeper) GetAllProfile(ctx sdk.Context) (list []types.Profile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Profile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
