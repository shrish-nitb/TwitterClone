package userservice

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

// CreateUser implements apiconnect.UserServiceHandler.
func (srv *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[api.CreateUserRequest],
) (*connect.Response[api.User], error) {
	entries := req.Msg.User
	if entries == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid format"))
	}
	if strings.Trim(entries.FirstName, " ") == "" || strings.Trim(entries.LastName, " ") == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide firstName and lastName"))
	}
	receipt, err := srv.User.Create().SetFirstName(req.Msg.User.FirstName).SetLastName(req.Msg.User.LastName).Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return buildUserResponse(receipt)
}
