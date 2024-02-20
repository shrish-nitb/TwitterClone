package userservice

import (
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

type UserService struct {
	User *ent.UserClient
}

func buildUserResponse(result *ent.User) (*connect.Response[api.User], error) {
	return connect.NewResponse(&api.User{
		Id:        strconv.Itoa(int(result.ID)),
		FirstName: result.FirstName,
		LastName:  result.LastName,
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
