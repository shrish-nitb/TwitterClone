package userservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/ent/user"
)

// ListUsers implements apiconnect.UserServiceHandler.
func (srv *UserService) ListUsers(ctx context.Context, req *connect.Request[api.ListUserRequest]) (*connect.Response[api.ListUserResponse], error) {
	pageSize := int(req.Msg.PageSize)
	pageToken, err := strconv.Atoi(req.Msg.PageToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("please provide valid page token"))
	}

	//getting the page via pageToken(ID)
	userList, err := srv.User.Query().Where(user.IDGTE(pageToken)).Limit(pageSize).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else if len(userList) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("nothing to show"))
	}

	//getting the nextPageToken if exist otherwise returning -1 if last page
	nextPageToken := -1
	nextPageUser, err := srv.User.Query().Where(user.IDGT(userList[len(userList)-1].ID)).Limit(1).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error while fetching users"))
	} else if len(nextPageUser) > 0 {
		nextPageToken = nextPageUser[0].ID
	}

	//buliding the response
	var response []*api.User
	for _, v := range userList {
		user := api.User{
			Id:        strconv.Itoa(int(v.ID)),
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
		response = append(response, &user)
	}
	return connect.NewResponse[api.ListUserResponse](&api.ListUserResponse{
		Users:         response,
		NextPageToken: strconv.Itoa(nextPageToken),
	}), nil
}
