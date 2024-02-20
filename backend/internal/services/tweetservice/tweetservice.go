package tweetservice

import (
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

type TweetService struct {
	Tweet *ent.TweetClient
}

func buildTweetResponse(result *ent.Tweet) (*connect.Response[api.Tweet], error) {
	return connect.NewResponse(&api.Tweet{
		Id:      strconv.Itoa(int(result.ID)),
		Content: result.Content,
		User:    result.User,
	}), nil
}

func checkQueryErr(err error) error {
	if err != nil {
		if ent.IsNotFound(err) {
			return connect.NewError(connect.CodeNotFound, errors.New("requested data not exist"))
		}
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}
