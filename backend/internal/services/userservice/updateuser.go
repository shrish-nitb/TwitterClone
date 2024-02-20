package userservice

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent"
)

// UpdateUser implements apiconnect.UserServiceHandler.
func (srv *UserService) UpdateUser(ctx context.Context, req *connect.Request[api.UpdateUserRequest]) (*connect.Response[api.User], error) {
	if req.Msg.FirstName == nil && req.Msg.LastName == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide a value to atleast one parameter"))
	}

	queryId, err := strconv.Atoi(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UserId invalid for updation"))
	}
	var receipt *ent.User

	if req.Msg.FirstName != nil && strings.Trim(*req.Msg.FirstName, " ") != "" {
		receipt, err = srv.User.UpdateOneID(queryId).SetFirstName(*req.Msg.FirstName).Save(ctx)
	}
	if err = checkQueryErr(err); err != nil {
		return nil, err
	}

	if req.Msg.LastName != nil && strings.Trim(*req.Msg.LastName, " ") != "" {
		receipt, err = srv.User.UpdateOneID(queryId).SetLastName(*req.Msg.LastName).Save(ctx)
	}
	if err = checkQueryErr(err); err != nil {
		return nil, err
	}

	if receipt != nil {
		return buildUserResponse(receipt)
	} else {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("fields cannot be blank"))
	}
}
