package userservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

// GetUser implements apiconnect.UserServiceHandler.
func (srv *UserService) GetUser(ctx context.Context, req *connect.Request[api.GetUserRequest]) (*connect.Response[api.User], error) {
	userID, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid userID"))
	}
	result, err := srv.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("requested data not exist"))
		}
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	return buildUserResponse(result)
}
