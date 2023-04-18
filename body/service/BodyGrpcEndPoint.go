package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	IdGrpcReq struct {
		Id int64
	}
	GetBodyByIdGrpcResp struct {
		Body interface{}
		// Id  int64
		Err string
	}
	BodyReq struct {
		Body Body
		Err  string
	}
	DefaultResp struct {
		Result string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById    endpoint.Endpoint
	CreateBody endpoint.Endpoint
	DeleteBody endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s IBodyService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById:    GrpcGetBodyByIdEndpoint(s),
		CreateBody: GrpcCreateBodyEndpoint(s),
		DeleteBody: GrpcDeleteBodyEndpoint(s),
	}
}

func GrpcGetBodyByIdEndpoint(s IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGrpcReq)
		bodyDetails, err := s.GetBodyById(ctx, int(req.Id))

		if err != nil {
			return GetBodyByIdGrpcResp{Body: nil, Err: "Id not found"}, nil
		}
		return GetBodyByIdGrpcResp{Body: bodyDetails, Err: ""}, nil
	}
}

func GrpcCreateBodyEndpoint(s IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BodyReq)
		result, err := s.CreateBody(ctx, req.Body)

		if err != nil {
			return DefaultResp{Result: "Create fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcDeleteBodyEndpoint(s IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGrpcReq)
		result, err := s.DeleteBody(ctx, int(req.Id))

		if err != nil {
			return DefaultResp{Result: "Delete fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}
