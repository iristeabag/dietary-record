package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	GetFoodByIdGrpcReq struct {
		Id int64
	}
	GetFoodByIdGrpcResp struct {
		Food interface{}
		// Id  int64
		Err string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s IFoodService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById: GetGrpcFoodByIdEndpoint(s),
	}
}

func GetGrpcFoodByIdEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoodByIdGrpcReq)
		foodDetails, err := s.GetFoodById(ctx, req.Id)

		if err != nil {
			return GetFoodByIdGrpcResp{Food: nil, Err: "Id not found"}, nil
		}
		// return GetFoodByIdGrpcResp{Id: req.Id, Err: ""}, nil
		return GetFoodByIdGrpcResp{Food: foodDetails, Err: ""}, nil
	}
}
