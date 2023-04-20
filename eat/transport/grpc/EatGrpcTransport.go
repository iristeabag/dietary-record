package service

import (
	"context"
	"errors"
	endpoint "go-kit-demo/eat/endpoint/grpc"
	svc "go-kit-demo/eat/service"
	pb "go-kit-demo/proto/eat"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcEatServer struct {
	getbyid   gt.Handler
	geteats   gt.Handler
	createeat gt.Handler
	updateeat gt.Handler
	deleteeat gt.Handler
}

func NewGRPCServer(endpoints endpoint.GrpcEndpoints, logger log.Logger) pb.EatServiceServer {
	return &grpcEatServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGrpcGetEatByIdRequest,
			EncodeGrpcResponse,
		),
		geteats: gt.NewServer(
			endpoints.GetEats,
			DecodeGrpcGetEatsequest,
			EncodeGrpcGetEatsResponse,
		),
		createeat: gt.NewServer(
			endpoints.CreateEat,
			DecodeGrpcCreateEatRequest,
			EncodeGrpcDefaultResponse,
		),
		updateeat: gt.NewServer(
			endpoints.UpdateEat,
			DecodeGrpcUpdateEatRequest,
			EncodeGrpcDefaultResponse,
		),
		deleteeat: gt.NewServer(
			endpoints.DeleteEat,
			DecodeGrpcDeleteEatRequest,
			EncodeGrpcDefaultResponse,
		),
	}
}

func (s *grpcEatServer) GetEatById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetEatByIdResponse, error) {
	_, resp, err := s.getbyid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetEatByIdResponse), nil
}

func (s *grpcEatServer) GetEats(ctx context.Context, req *pb.GetEatsRequest) (*pb.GetEatsResponse, error) {
	_, resp, err := s.geteats.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetEatsResponse), nil
}

func (s *grpcEatServer) CreateEat(ctx context.Context, req *pb.EatRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.createeat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcEatServer) UpdateEat(ctx context.Context, req *pb.EatRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.updateeat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcEatServer) DeleteEat(ctx context.Context, req *pb.GetByIdRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.deleteeat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func DecodeGrpcGetEatByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetByIdRequest)
	return endpoint.GetEatByIdGrpcReq{Id: req.Id}, nil
}

func DecodeGrpcGetEatsequest(_ context.Context, request interface{}) (interface{}, error) {
	// request.(*pb.GetEatsRequest)
	return endpoint.GetEatsReq{}, nil
}

func EncodeGrpcResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetEatByIdGrpcResp)

	if resp.Eat == nil {
		return nil, errors.New("id not found")
	}

	eat := resp.Eat.(svc.Eat)

	result := pb.Eat{
		Eatid:   eat.Id,
		Foodid:  eat.Foodid,
		Name:    eat.Name,
		Amount:  float32(eat.Amount),
		Unit:    eat.Unit,
		Carb:    float32(eat.Carb),
		Portein: float32(eat.Portein),
		Fat:     float32(eat.Fat),
		Cal:     float32(eat.Cal),
	}

	return &pb.GetEatByIdResponse{
		Eat: &result,
	}, nil

}

func EncodeGrpcGetEatsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetEatsGrpcResp)

	if resp.Eat == nil {
		return nil, errors.New("id not found")
	}

	eatlist := resp.Eat.([]interface{})

	var results []*pb.Eat
	for _, item := range eatlist {
		eat := item.(svc.Eat)
		result := pb.Eat{
			Eatid:   eat.Id,
			Foodid:  eat.Foodid,
			Name:    eat.Name,
			Amount:  float32(eat.Amount),
			Unit:    eat.Unit,
			Carb:    float32(eat.Carb),
			Portein: float32(eat.Portein),
			Fat:     float32(eat.Fat),
			Cal:     float32(eat.Cal),
		}
		results = append(results, &result)
	}

	return &pb.GetEatsResponse{
		Eat: results,
	}, nil
}

func DecodeGrpcCreateEatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.EatRequest)

	eat := svc.Eat{
		Foodid:  req.Eat.Foodid,
		Name:    req.Eat.Name,
		Amount:  float64(req.Eat.Amount),
		Unit:    req.Eat.Unit,
		Carb:    float64(req.Eat.Carb),
		Portein: float64(req.Eat.Portein),
		Fat:     float64(req.Eat.Fat),
		Cal:     float64(req.Eat.Cal),
	}

	return endpoint.EatReq{Eat: eat}, nil
}

func DecodeGrpcUpdateEatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.EatRequest)
	eat := svc.Eat{
		Id:      req.Eat.Eatid,
		Foodid:  req.Eat.Foodid,
		Name:    req.Eat.Name,
		Amount:  float64(req.Eat.Amount),
		Unit:    req.Eat.Unit,
		Carb:    float64(req.Eat.Carb),
		Portein: float64(req.Eat.Portein),
		Fat:     float64(req.Eat.Fat),
		Cal:     float64(req.Eat.Cal),
	}
	return endpoint.EatReq{Eat: eat}, nil
}

func DecodeGrpcDeleteEatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetByIdRequest)
	return endpoint.GetEatByIdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcDefaultResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DefaultResp)

	return &pb.DefaultResponse{Result: resp.Result}, nil

}
