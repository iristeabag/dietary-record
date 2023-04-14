package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	GetFoodByIdRequest struct {
		Id string `json:"foodid"`
	}
	GetFoodByIdResponse struct {
		Food interface{} `json:"food,omitempty"`
		Err  string      `json:"error,omitempty"`
	}
	GetAllFoodsRequest struct{}

	GetAllFoodsResponse struct {
		Food interface{} `json:"food,omitempty"`
		Err  string      `json:"error,omitempty"`
	}
	CreateFoodRequest struct {
		food Food
	}
	CreateFoodResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	UpdateFoodRequest struct {
		food Food
	}
	UpdateFoodResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	DeleteFoodRequest struct {
		Foodid string `json:"foodid"`
	}

	DeleteFoodResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
)

func GetFoodByIdEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFoodByIdRequest)
		foodDetails, err := s.GetFoodById(ctx, req.Id)
		if err != nil {
			return GetFoodByIdResponse{Food: foodDetails, Err: "Id not found"}, nil
		}
		return GetFoodByIdResponse{Food: foodDetails, Err: ""}, nil
	}
}
func GetAllFoodsEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		foodDetails, err := s.GetAllFoods(ctx)
		if err != nil {
			return GetAllFoodsResponse{Food: foodDetails, Err: "no data found"}, nil
		}
		return GetAllFoodsResponse{Food: foodDetails, Err: ""}, nil
	}
}

func CreateFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFoodRequest)
		msg, err := s.CreateFood(ctx, req.food)
		return CreateFoodResponse{Msg: msg, Err: err}, nil
	}
}

func UpdateFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateFoodRequest)
		msg, err := s.UpdateFood(ctx, req.food)
		return msg, err
	}
}

func DeleteFoodEndpoint(s IFoodService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFoodRequest)
		msg, err := s.DeleteFood(ctx, req.Foodid)
		if err != nil {
			return DeleteFoodResponse{Msg: msg, Err: err}, nil
		}
		return DeleteFoodResponse{Msg: msg, Err: nil}, nil
	}
}
