package userservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteUser implements apiconnect.UserServiceHandler.
func (srv *UserService) DeleteUser(ctx context.Context, req *connect.Request[api.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	userID, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid id param"))
	}
	if err = srv.User.DeleteOneID(userID).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
