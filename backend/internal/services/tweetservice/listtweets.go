package tweetservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent/tweet"
)

// ListTweets implements apiconnect.TweetServiceHandler.
func (srv *TweetService) ListTweets(ctx context.Context, req *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error) {
	pageSize := int(req.Msg.PageSize)
	pageToken, err := strconv.Atoi(req.Msg.PageToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid page token"))
	}

	_, err = strconv.Atoi(req.Msg.User)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid userID"))
	}

	//getting the page via pageToken(ID)
	tweetList, err := srv.Tweet.Query().Where(tweet.IDGTE(pageToken), tweet.User(req.Msg.User)).Limit(pageSize).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else if len(tweetList) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("nothing to show"))
	}

	//getting the nextPageToken if exist otherwise returning -1 if last page
	nextPageToken := -1
	nextPageUser, err := srv.Tweet.Query().Where(tweet.IDGT(tweetList[len(tweetList)-1].ID)).Limit(1).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error while fetching users"))
	} else if len(nextPageUser) > 0 {
		nextPageToken = nextPageUser[0].ID
	}

	//buliding the response
	var response []*api.Tweet
	for _, v := range tweetList {
		tweet := api.Tweet{
			Id:      strconv.Itoa(int(v.ID)),
			Content: v.Content,
			User:    v.User,
		}
		response = append(response, &tweet)
	}
	return connect.NewResponse[api.ListTweetResponse](&api.ListTweetResponse{
		Tweets:        response,
		NextPageToken: strconv.Itoa(nextPageToken),
	}), nil
}
