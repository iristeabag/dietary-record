package endpoint

import (
	"context"

	svc "go-kit-demo/body/service"

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
	GetBodysReq      struct{}
	GetBodysGrpcResp struct {
		Body interface{}
		Err  string
	}
	BodyReq struct {
		Body svc.Body
		Err  string
	}
	DefaultResp struct {
		Result string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById    endpoint.Endpoint
	GetBodys   endpoint.Endpoint
	CreateBody endpoint.Endpoint
	UpdateBody endpoint.Endpoint
	DeleteBody endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s svc.IBodyService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById:    GrpcGetBodyByIdEndpoint(s),
		GetBodys:   GrpcGetBodysEndpoint(s),
		CreateBody: GrpcCreateBodyEndpoint(s),
		UpdateBody: GrpcUpdateBodyEndpoint(s),
		DeleteBody: GrpcDeleteBodyEndpoint(s),
	}
}

func GrpcGetBodyByIdEndpoint(s svc.IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGrpcReq)
		bodyDetails, err := s.GetBodyById(ctx, int(req.Id))

		if err != nil {
			return GetBodyByIdGrpcResp{Body: nil, Err: "Id not found"}, nil
		}
		return GetBodyByIdGrpcResp{Body: bodyDetails, Err: ""}, nil
	}
}

func GrpcGetBodysEndpoint(s svc.IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		bodyDetails, err := s.GetAllBodys(ctx)

		if err != nil {
			return GetBodysGrpcResp{Body: nil, Err: "Id not found"}, nil
		}
		return GetBodysGrpcResp{Body: bodyDetails, Err: ""}, nil
	}
}

func GrpcCreateBodyEndpoint(s svc.IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BodyReq)
		result, err := s.CreateBody(ctx, req.Body)

		if err != nil {
			return DefaultResp{Result: "Create fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcUpdateBodyEndpoint(s svc.IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BodyReq)
		result, err := s.UpdateBody(ctx, req.Body)

		if err != nil {
			return DefaultResp{Result: "Update fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcDeleteBodyEndpoint(s svc.IBodyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGrpcReq)
		result, err := s.DeleteBody(ctx, int(req.Id))

		if err != nil {
			return DefaultResp{Result: "Delete fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}
