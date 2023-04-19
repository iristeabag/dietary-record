package service

import (
	"context"
	svc "go-kit-demo/food/service"

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
	GetFoodsReq      struct{}
	GetFoodsGrpcResp struct {
		Food interface{}
		Err  string
	}
	FoodReq struct {
		Food svc.Food
		Err  string
	}
	DefaultResp struct {
		Result string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById    endpoint.Endpoint
	GetFoods   endpoint.Endpoint
	CreateFood endpoint.Endpoint
	UpdateFood endpoint.Endpoint
	DeleteFood endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s svc.IFoodService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById:    GrpcGetFoodByIdEndpoint(s),
		GetFoods:   GrpcGetFoodsEndpoint(s),
		CreateFood: GrpcCreateFoodEndpoint(s),
		UpdateFood: GrpcUpdateFoodEndpoint(s),
		DeleteFood: GrpcDeleteFoodEndpoint(s),
	}
}

func GrpcGetFoodByIdEndpoint(s svc.IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoodByIdGrpcReq)
		foodDetails, err := s.GetFoodById(ctx, int(req.Id))

		if err != nil {
			return GetFoodByIdGrpcResp{Food: nil, Err: "Id not found"}, nil
		}
		return GetFoodByIdGrpcResp{Food: foodDetails, Err: ""}, nil
	}
}

func GrpcGetFoodsEndpoint(s svc.IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		bodyDetails, err := s.GetAllFoods(ctx)

		if err != nil {
			return GetFoodsGrpcResp{Food: nil, Err: "Id not found"}, nil
		}
		return GetFoodsGrpcResp{Food: bodyDetails, Err: ""}, nil
	}
}

func GrpcCreateFoodEndpoint(s svc.IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FoodReq)
		result, err := s.CreateFood(ctx, req.Food)

		if err != nil {
			return DefaultResp{Result: "Create fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcUpdateFoodEndpoint(s svc.IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FoodReq)
		result, err := s.UpdateFood(ctx, req.Food)

		if err != nil {
			return DefaultResp{Result: "Update fail"}, err
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcDeleteFoodEndpoint(s svc.IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoodByIdGrpcReq)
		result, err := s.DeleteFood(ctx, int(req.Id))

		if err != nil {
			return DefaultResp{Result: "Delete fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}
