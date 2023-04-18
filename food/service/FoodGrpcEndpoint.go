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
	FoodReq struct {
		Food Food
		Err  string
	}
	DefaultResp struct {
		Result string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById    endpoint.Endpoint
	CreateFood endpoint.Endpoint
	UpdateFood endpoint.Endpoint
	DeleteFood endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s IFoodService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById:    GrpcGetFoodByIdEndpoint(s),
		CreateFood: GrpcCreateFoodEndpoint(s),
		UpdateFood: GrpcUpdateFoodEndpoint(s),
		DeleteFood: GrpcDeleteFoodEndpoint(s),
	}
}

func GrpcGetFoodByIdEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoodByIdGrpcReq)
		foodDetails, err := s.GetFoodById(ctx, req.Id)

		if err != nil {
			return GetFoodByIdGrpcResp{Food: nil, Err: "Id not found"}, nil
		}
		return GetFoodByIdGrpcResp{Food: foodDetails, Err: ""}, nil
	}
}

func GrpcCreateFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FoodReq)
		result, err := s.CreateFood(ctx, req.Food)

		if err != nil {
			return DefaultResp{Result: "Create fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcUpdateFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FoodReq)
		result, err := s.UpdateFood(ctx, req.Food.Foodid, req.Food)

		if err != nil {
			return DefaultResp{Result: "Update fail"}, err
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcDeleteFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoodByIdGrpcReq)
		result, err := s.DeleteFood(ctx, int(req.Id))

		if err != nil {
			return DefaultResp{Result: "Delete fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}
