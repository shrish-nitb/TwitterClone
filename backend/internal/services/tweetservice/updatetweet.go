package tweetservice

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

// UpdateTweet implements apiconnect.TweetServiceHandler.
func (srv *TweetService) UpdateTweet(ctx context.Context, req *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error) {
	if req.Msg.Content == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide content value"))
	}

	queryId, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("TweetId invalid for updation"))
	}

	var receipt *ent.Tweet

	if req.Msg.Content != nil && strings.Trim(*req.Msg.Content, " ") != "" {
		receipt, err = srv.Tweet.UpdateOneID(queryId).SetContent(*req.Msg.Content).Save(ctx)
	}
	if err = checkQueryErr(err); err != nil {
		return nil, err
	}

	if receipt != nil {
		return buildTweetResponse(receipt)
	} else {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("fields cannot be blank"))
	}

}
