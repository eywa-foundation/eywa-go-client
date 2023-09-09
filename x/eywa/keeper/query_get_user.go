package keeper

import (
	"context"

	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetUser(goCtx context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	user, found := k.GetUserByAddress(ctx, req.Submitter)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetUserResponse{User: &user}, nil
}
