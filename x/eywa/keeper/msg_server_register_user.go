package keeper

import (
	"context"

	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterUser(goCtx context.Context, msg *types.MsgRegisterUser) (*types.MsgRegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var user = types.User{
		Submitter: msg.Creator,
		Pubkey:    msg.Pubkey,
	}

	id := k.CreateUser(ctx, user)

	return &types.MsgRegisterUserResponse{Id: id}, nil
}
