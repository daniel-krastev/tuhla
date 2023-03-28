package controller

import (
	"context"
	"fmt"

	"tuhla/services/grpc_gateway/proto/houses/housespb"
	"tuhla/services/houses/proto/housesmodelpb"
	"tuhla/services/houses/proto/housesservicepb"
)

func (c *Controller) CreateHouse(ctx context.Context, createHouseReq *housespb.CreateHouseRequest) (*housespb.House, error) {
	house, err := c.housesService.CreateHouse(
		ctx,
		&housesservicepb.CreateHouseRequest{
			RequestId: createHouseReq.RequestId,
			House: &housesmodelpb.House{
				Id: createHouseReq.House.Id,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error from internal houses service: %v", err)
	}
	return &housespb.House{Id: fmt.Sprintf("%s_from.GRPCGateway", house.Id)}, nil
}
