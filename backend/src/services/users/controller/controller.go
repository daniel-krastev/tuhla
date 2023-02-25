package controller

import (
	"context"
	"errors"
	"tuhla/services/users/proto/usersmodelpb"
	"tuhla/services/users/proto/usersservicepb"
)

type Storage interface {
	CreateUser(requestID, userID string) (string, error)
}

type Controller struct {
	storage Storage
	usersservicepb.UnimplementedUsersServer
}

func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

func (c *Controller) CreateUser(ctx context.Context, createUserReq *usersservicepb.CreateUserRequest) (*usersmodelpb.User, error) {
	if createUserReq == nil || createUserReq.User == nil {
		return nil, errors.New("no request found")
	}
	createdUser := &usersmodelpb.User{
		Id: createUserReq.User.Id,
	}
	return createdUser, nil
}
