package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	GetBodyByIdRequest struct {
		Id string `json:"bodyid"`
	}
	GetBodyByIdResponse struct {
		Body interface{} `json:"body,omitempty"`
		Err  string      `json:"error,omitempty"`
	}

	GetAllBodysRequest struct {
		Date string `json:"date"`
	}

	GetAllBodysResponse struct {
		Body interface{} `json:"body,omitempty"`
		Err  string      `json:"error,omitempty"`
	}
	CreateBodyRequest struct {
		body Body
	}
	CreateBodyResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	UpdateBodyRequest struct {
		body Body
	}
	UpdateBodyResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	DeleteBodyRequest struct {
		Id string `json:"bodyid"`
	}

	DeleteBodyResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
)

func GetBodyByIdEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBodyByIdRequest)
		bodyDetails, err := e.GetBodyById(ctx, req.Id)
		if err != nil {
			return GetBodyByIdResponse{Body: bodyDetails, Err: "Id not found"}, nil
		}
		return GetBodyByIdResponse{Body: bodyDetails, Err: ""}, nil
	}
}

func GetAllBodysEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllBodysRequest)
		var bodyDetails interface{}
		var err error
		if req.Date != "" {
			bodyDetails, err = e.GetBodyByDate(ctx, req.Date)
		} else {
			bodyDetails, err = e.GetAllBodys(ctx)
		}

		if err != nil {
			return GetAllBodysResponse{Body: bodyDetails, Err: "no data found"}, nil
		}
		return GetAllBodysResponse{Body: bodyDetails, Err: ""}, nil
	}
}

func CreateBodyEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBodyRequest)
		msg, err := e.CreateBody(ctx, req.body)
		return CreateBodyResponse{Msg: msg, Err: err}, nil
	}
}

func UpdateBodyEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBodyRequest)
		msg, err := e.UpdateBody(ctx, req.body)
		return msg, err
	}
}

func DeleteBodyEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBodyRequest)
		msg, err := e.DeleteBody(ctx, req.Id)
		if err != nil {
			return DeleteBodyResponse{Msg: msg, Err: err}, nil
		}
		return DeleteBodyResponse{Msg: msg, Err: nil}, nil
	}
}
