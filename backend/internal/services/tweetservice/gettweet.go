package tweetservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

// GetTweet implements apiconnect.TweetServiceHandler.
func (srv *TweetService) GetTweet(ctx context.Context, req *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error) {
	tweetID, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid tweetID"))
	}
	result, err := srv.Tweet.Get(ctx, tweetID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("requested data not exist"))
		}
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	return buildTweetResponse(result)
}
