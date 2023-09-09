package keeper

import (
	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateUser(ctx sdk.Context, user types.User) uint64 {
	user.Id = 0
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	appendedValue := k.cdc.MustMarshal(&user)
	store.Set(GetUserAddressBytes(user.Submitter), appendedValue)
	return user.Id
}

func (k Keeper) GetUserByAddress(ctx sdk.Context, submitter string) (val types.User, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))

	b := store.Get(GetUserAddressBytes(submitter))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func GetUserAddressBytes(submitter string) []byte {
	return []byte(submitter)
}
