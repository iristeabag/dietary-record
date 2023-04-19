package service

import (
	"context"
	"strconv"

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

	GetAllBodysRequest  struct{}
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
		Id   string `json:"bodyid"`
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
		id, _ := strconv.Atoi(req.Id)
		bodyDetails, err := e.GetBodyById(ctx, id)
		if err != nil {
			return GetBodyByIdResponse{Body: bodyDetails, Err: "Id not found"}, nil
		}
		return GetBodyByIdResponse{Body: bodyDetails, Err: ""}, nil
	}
}

func GetAllBodysEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		bodyDetails, err := e.GetAllBodys(ctx)

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

func UpdateBodyEndpoint(s IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBodyRequest)
		req.body.Id = req.Id
		msg, err := s.UpdateBody(ctx, req.body)
		return UpdateBodyResponse{Msg: msg, Err: nil}, err
	}
}

func DeleteBodyEndpoint(e IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBodyRequest)
		id, _ := strconv.Atoi(req.Id)
		msg, err := e.DeleteBody(ctx, id)
		if err != nil {
			return DeleteBodyResponse{Msg: msg, Err: err}, nil
		}
		return DeleteBodyResponse{Msg: msg, Err: nil}, nil
	}
}
