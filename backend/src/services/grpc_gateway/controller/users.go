package controller

import (
	"context"
	"fmt"

	"tuhla/services/grpc_gateway/proto/users/userspb"
	"tuhla/services/users/proto/usersmodelpb"
	"tuhla/services/users/proto/usersservicepb"
)

func (c *Controller) CreateUser(ctx context.Context, createUserReq *userspb.CreateUserRequest) (*userspb.User, error) {
	user, err := c.userService.CreateUser(
		ctx,
		&usersservicepb.CreateUserRequest{
			RequestId: createUserReq.RequestId,
			User: &usersmodelpb.User{
				Id: createUserReq.User.Id,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error from internal users service: %v", err)
	}
	return &userspb.User{Id: fmt.Sprintf("%s_from.GRPCGateway", user.Id)}, nil
}
