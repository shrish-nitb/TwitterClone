package twitterapp_test

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"testing"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
)

func TestTwitterApp(t *testing.T) {
	// go twitterapp.Run()
	userClient := apiconnect.NewUserServiceClient(http.DefaultClient, "http://127.0.0.1:8080")

	//testing CreateUser()
	createUserTest, err := userClient.CreateUser(context.Background(), connect.NewRequest(&api.CreateUserRequest{
		User: &api.User{
			FirstName: "hello",
			LastName:  "bar",
		},
	}))
	ok(t, err)
	notEmpty(t, createUserTest.Msg.Id)

	//testing GetUser()
	getUserTest, err := userClient.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{
		Id: createUserTest.Msg.Id,
	}))
	ok(t, err)
	equalStr(t, getUserTest.Msg.FirstName, createUserTest.Msg.FirstName)
	equalStr(t, getUserTest.Msg.LastName, createUserTest.Msg.LastName)

	// testing ListUsers()
	listUserTest, err := userClient.ListUsers(context.Background(), connect.NewRequest(&api.ListUserRequest{
		PageSize:  1,
		PageToken: createUserTest.Msg.Id,
	}))
	ok(t, err)
	equalStr(t, listUserTest.Msg.NextPageToken, "-1")

	// testing UpdateUser()
	updateUserTest, err := userClient.UpdateUser(context.Background(), connect.NewRequest(&api.UpdateUserRequest{
		Id:        createUserTest.Msg.Id,
		FirstName: &createUserTest.Msg.LastName,
		LastName:  &createUserTest.Msg.FirstName,
	}))

	ok(t, err)
	equalStr(t, updateUserTest.Msg.FirstName, createUserTest.Msg.LastName)
	equalStr(t, updateUserTest.Msg.LastName, createUserTest.Msg.FirstName)

	//trying to update single entry
	updateUserTest, err = userClient.UpdateUser(context.Background(), connect.NewRequest(&api.UpdateUserRequest{
		Id:        createUserTest.Msg.Id,
		FirstName: &createUserTest.Msg.FirstName,
		LastName:  nil,
	}))
	ok(t, err)
	equalStr(t, updateUserTest.Msg.FirstName, createUserTest.Msg.FirstName)
	equalStr(t, updateUserTest.Msg.LastName, createUserTest.Msg.FirstName)

	tweetClient := apiconnect.NewTweetServiceClient(http.DefaultClient, "http://127.0.0.1:8080")
	//testing CreateTweet()
	createTweetTest, err := tweetClient.CreateTweet(context.Background(), connect.NewRequest(&api.CreateTweetRequest{
		Tweet: &api.Tweet{
			Content: "new tweet2",
			User:    createUserTest.Msg.Id,
		},
	}))
	ok(t, err)
	notEmpty(t, createTweetTest.Msg.Id)

	//testing GetTweet()
	getTweetTest, err := tweetClient.GetTweet(context.Background(), connect.NewRequest(&api.GetTweetRequest{
		Id: createTweetTest.Msg.Id,
	}))
	ok(t, err)
	equalStr(t, getTweetTest.Msg.Content, createTweetTest.Msg.Content)

	//testing ListTweets()
	createdId, _ := strconv.Atoi(createTweetTest.Msg.Id)
	listTweetTest, err := tweetClient.ListTweets(context.Background(), connect.NewRequest(&api.ListTweetRequest{
		PageSize:  int32(createdId),
		PageToken: "1",
		User:      createUserTest.Msg.Id,
	}))

	if createdId != len(listTweetTest.Msg.Tweets) && err != nil {
		t.FailNow()
	}

	//testing UpdateTweet()
	var newcontent string = "changed content"
	updateTweetTest, err := tweetClient.UpdateTweet(context.Background(), connect.NewRequest(&api.UpdateTweetRequest{
		Id:      createTweetTest.Msg.Id,
		Content: &newcontent,
	}))

	ok(t, err)
	equalStr(t, updateTweetTest.Msg.Content, "changed content")

	//testing DeleteTweet()
	_, err = tweetClient.DeleteTweet(context.Background(), connect.NewRequest(&api.DeleteTweetRequest{
		Id: createTweetTest.Msg.Id,
	}))

	ok(t, err)

	// testing DeleteUser()
	_, err = userClient.DeleteUser(context.Background(), connect.NewRequest(&api.DeleteUserRequest{
		Id: createUserTest.Msg.Id,
	}))
	ok(t, err)

	//testing GetUser() again against non-existant ID
	if _, err = userClient.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{
		Id: createUserTest.Msg.Id,
	})); err == nil {
		t.FailNow()
	}

	//testing UpdateUser() for non-existant ID
	if _, err = userClient.UpdateUser(context.Background(), connect.NewRequest(&api.UpdateUserRequest{
		Id: createUserTest.Msg.Id,
	})); err == nil {
		t.FailNow()
	}

	//testing UpdateTweet() with non-existant tweet ID
	if _, err = tweetClient.UpdateTweet(context.Background(), connect.NewRequest(&api.UpdateTweetRequest{
		Id: createTweetTest.Msg.Id,
	})); err == nil {
		t.FailNow()
	}

}

func ok(t *testing.T, err error) {
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

func notEmpty(t *testing.T, v string) {
	if v == "" {
		t.FailNow()
	}
}

func equalStr(t *testing.T, f, s string) {
	if f != s {
		t.FailNow()
	}
}
