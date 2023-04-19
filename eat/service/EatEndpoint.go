package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

type (
	GetEatByIdRequest struct {
		Id string `json:"eatid"`
	}
	GetEatByIdResponse struct {
		Eat interface{} `json:"eat,omitempty"`
		Err string      `json:"error,omitempty"`
	}

	GetTotalEatByDateRequest struct {
		Date string `json:"date"`
	}
	GetTotalEatByDateResponse struct {
		TotalEat interface{} `json:"totaleat,omitempty"`
		Err      string      `json:"error,omitempty"`
	}

	GetAllEatsRequest struct {
		Date string `json:"date"`
	}

	GetAllEatsResponse struct {
		Eat interface{} `json:"eat,omitempty"`
		Err string      `json:"error,omitempty"`
	}
	CreateEatRequest struct {
		eat Eat
	}
	CreateEatResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	UpdateEatRequest struct {
		Id  string `json:"eatid"`
		eat Eat
	}
	UpdateEatResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	DeleteEatRequest struct {
		Id string `json:"eatid"`
	}

	DeleteEatResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
)

func GetEatByIdEndpoint(e IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEatByIdRequest)
		id, _ := strconv.Atoi(req.Id)
		eatDetails, err := e.GetEatById(ctx, int(id))
		if err != nil {
			return GetEatByIdResponse{Eat: eatDetails, Err: "Id not found"}, nil
		}
		return GetEatByIdResponse{Eat: eatDetails, Err: ""}, nil
	}
}

func GetAllEatsEndpoint(e IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllEatsRequest)
		var eatDetails interface{}
		var err error
		if req.Date != "" {
			eatDetails, err = e.GetEatsByDate(ctx, req.Date)
		} else {
			eatDetails, err = e.GetAllEats(ctx)
		}

		if err != nil {
			return GetAllEatsResponse{Eat: eatDetails, Err: "no data found"}, nil
		}
		return GetAllEatsResponse{Eat: eatDetails, Err: ""}, nil
	}
}

func CreateEatEndpoint(e IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEatRequest)
		msg, err := e.CreateEat(ctx, req.eat)
		return CreateEatResponse{Msg: msg, Err: err}, nil
	}
}

func UpdateEatEndpoint(e IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEatRequest)
		msg, err := e.UpdateEat(ctx, req.eat)
		return UpdateEatResponse{Msg: msg, Err: nil}, err
	}
}

func DeleteEatEndpoint(e IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEatRequest)
		id, _ := strconv.Atoi(req.Id)
		msg, err := e.DeleteEat(ctx, int(id))
		if err != nil {
			return DeleteEatResponse{Msg: msg, Err: err}, nil
		}
		return DeleteEatResponse{Msg: msg, Err: nil}, nil
	}
}
