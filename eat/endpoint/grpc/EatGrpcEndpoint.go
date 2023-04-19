package service

import (
	"context"

	svc "go-kit-demo/eat/service"

	"github.com/go-kit/kit/endpoint"
)

type (
	GetEatByIdGrpcReq struct {
		Id int64
	}
	GetEatByIdGrpcResp struct {
		Eat interface{}
		// Id  int64
		Err string
	}
	GetEatsReq      struct{}
	GetEatsGrpcResp struct {
		Eat interface{}
		Err string
	}
	EatReq struct {
		Eat svc.Eat
		Err string
	}
	DefaultResp struct {
		Result string
	}
)

// Endpoints struct holds the list of endpoints definition
type GrpcEndpoints struct {
	GetById   endpoint.Endpoint
	GetEats   endpoint.Endpoint
	CreateEat endpoint.Endpoint
	UpdateEat endpoint.Endpoint
	DeleteEat endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeGrpcEndpoints(s svc.IEatService) GrpcEndpoints {
	return GrpcEndpoints{
		GetById:   GrpcGetEatByIdEndpoint(s),
		GetEats:   GrpcGetEatsEndpoint(s),
		CreateEat: GrpcCreateEatEndpoint(s),
		UpdateEat: GrpcUpdateEatEndpoint(s),
		DeleteEat: GrpcDeleteEatEndpoint(s),
	}
}

func GrpcGetEatByIdEndpoint(s svc.IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetEatByIdGrpcReq)
		eatDetails, err := s.GetEatById(ctx, int(req.Id))

		if err != nil {
			return GetEatByIdGrpcResp{Eat: nil, Err: "Id not found"}, nil
		}
		return GetEatByIdGrpcResp{Eat: eatDetails, Err: ""}, nil
	}
}

func GrpcGetEatsEndpoint(s svc.IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		bodyDetails, err := s.GetAllEats(ctx)

		if err != nil {
			return GetEatsGrpcResp{Eat: nil, Err: "Id not found"}, nil
		}
		return GetEatsGrpcResp{Eat: bodyDetails, Err: ""}, nil
	}
}

func GrpcCreateEatEndpoint(s svc.IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EatReq)
		result, err := s.CreateEat(ctx, req.Eat)

		if err != nil {
			return DefaultResp{Result: "Create fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcUpdateEatEndpoint(s svc.IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EatReq)
		result, err := s.UpdateEat(ctx, req.Eat)

		if err != nil {
			return DefaultResp{Result: "Update fail"}, err
		}
		return DefaultResp{Result: result}, nil
	}
}

func GrpcDeleteEatEndpoint(s svc.IEatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetEatByIdGrpcReq)
		result, err := s.DeleteEat(ctx, int(req.Id))

		if err != nil {
			return DefaultResp{Result: "Delete fail"}, nil
		}
		return DefaultResp{Result: result}, nil
	}
}
