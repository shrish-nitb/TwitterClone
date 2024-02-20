package services

import (
	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
	"github.com/codesmith-dev/twitter/internal/services/tweetservice"
	"github.com/codesmith-dev/twitter/internal/services/userservice"
)

func NewUserServiceHandler(UserClient *ent.UserClient) apiconnect.UserServiceHandler {
	return &userservice.UserService{
		User: UserClient,
	}
}

func NewTweetServiceHandler(TweetClient *ent.TweetClient) apiconnect.TweetServiceHandler {
	return &tweetservice.TweetService{
		Tweet: TweetClient,
	}
}
