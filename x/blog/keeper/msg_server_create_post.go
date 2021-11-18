package keeper

import (
	"context"

	"github.com/cosmonaut/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type post 

	var post = types.Post {
		Creator : msg.Creator,
		Time: msg.Time,
		Title: msg.Title,
		Body: msg.Body,
	}

	id := k.AppendPost(ctx,post)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatePostResponse{id }, nil

}
