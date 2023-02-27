package controller

import (
	"tuhla/services/houses/proto/housesservicepb"
	"tuhla/services/users/proto/usersservicepb"

	"tuhla/services/grpc_gateway/proto/houses/housespb"
	"tuhla/services/grpc_gateway/proto/users/userspb"
)

type Controller struct {
	userService   usersservicepb.UsersClient
	housesService housesservicepb.HousesClient

	userspb.UnimplementedUsersServer
	housespb.UnimplementedHousesServer
}

func New(userService usersservicepb.UsersClient, housesService housesservicepb.HousesClient) *Controller {
	return &Controller{
		userService:   userService,
		housesService: housesService,
	}
}
