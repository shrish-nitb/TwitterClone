package tweetservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteTweet implements apiconnect.TweetServiceHandler.
func (srv *TweetService) DeleteTweet(ctx context.Context, req *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	tweetID, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid id param"))
	}
	if err = srv.Tweet.DeleteOneID(tweetID).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
