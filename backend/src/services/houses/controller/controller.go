package controller

import (
	"context"
	"errors"
	"fmt"

	"tuhla/services/houses/proto/housesmodelpb"
	"tuhla/services/houses/proto/housesservicepb"
)

type Controller struct {
	housesservicepb.UnimplementedHousesServer
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) CreateHouse(ctx context.Context, createHousesReq *housesservicepb.CreateHouseRequest) (*housesmodelpb.House, error) {
	if createHousesReq == nil || createHousesReq.House == nil {
		return nil, errors.New("no request found")
	}
	createdHouse := &housesmodelpb.House{
		Id: fmt.Sprintf("%s_fm.HousesService", createHousesReq.House.Id),
	}
	return createdHouse, nil
}
