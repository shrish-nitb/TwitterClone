package tweetservice

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

// CreateTweet implements apiconnect.TweetServiceHandler.
func (srv *TweetService) CreateTweet(ctx context.Context, req *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error) {
	entries := req.Msg.Tweet
	if entries == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid format"))
	}
	if strings.Trim(entries.Content, " ") == "" || strings.Trim(entries.User, " ") == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide Content and User"))
	}
	receipt, err := srv.Tweet.Create().SetContent(req.Msg.Tweet.Content).SetUser(req.Msg.Tweet.User).Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return buildTweetResponse(receipt)
}
