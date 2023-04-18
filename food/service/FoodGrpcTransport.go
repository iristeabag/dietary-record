package service

import (
	"context"
	"errors"
	pb "go-kit-demo/food/proto"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcFoodServer struct {
	getbyid    gt.Handler
	createfood gt.Handler
	updatefood gt.Handler
	deletefood gt.Handler
}

func NewGRPCServer(endpoints GrpcEndpoints, logger log.Logger) pb.FoodServiceServer {
	return &grpcFoodServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGrpcGetFoodByIdRequest,
			EncodeGrpcResponse,
		),
		createfood: gt.NewServer(
			endpoints.CreateFood,
			DecodeGrpcCreateFoodRequest,
			EncodeGrpcDefaultResponse,
		),
		updatefood: gt.NewServer(
			endpoints.UpdateFood,
			DecodeGrpcUpdateFoodRequest,
			EncodeGrpcDefaultResponse,
		),
		deletefood: gt.NewServer(
			endpoints.DeleteFood,
			DecodeGrpcDeleteFoodRequest,
			EncodeGrpcDefaultResponse,
		),
	}
}

func (s *grpcFoodServer) GetFoodById(ctx context.Context, req *pb.GetFoodByIdRequest) (*pb.GetFoodByIdResponse, error) {
	_, resp, err := s.getbyid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetFoodByIdResponse), nil
}

func (s *grpcFoodServer) CreateFood(ctx context.Context, req *pb.CreateFoodRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.createfood.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcFoodServer) UpdateFood(ctx context.Context, req *pb.UpdateFoodRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.updatefood.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcFoodServer) DeleteFood(ctx context.Context, req *pb.DeleteFoodRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.deletefood.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func DecodeGrpcGetFoodByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetFoodByIdRequest)
	return GetFoodByIdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetFoodByIdGrpcResp)

	if resp.Food == nil {
		return nil, errors.New("id not found")
	}

	food := resp.Food.(Food)
	return &pb.GetFoodByIdResponse{
		Foodid:  food.Foodid,
		Name:    food.Name,
		Brand:   food.Brand,
		Amount:  float32(food.Amount),
		Unit:    food.Unit,
		Carb:    float32(food.Carb),
		Portein: float32(food.Portein),
		Fat:     float32(food.Fat),
		Cal:     float32(food.Cal),
	}, nil

}

func DecodeGrpcCreateFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateFoodRequest)

	food := Food{
		Name:    req.Name,
		Brand:   req.Brand,
		Amount:  float64(req.Amount),
		Unit:    req.Unit,
		Carb:    float64(req.Carb),
		Portein: float64(req.Portein),
		Fat:     float64(req.Fat),
		Cal:     float64(req.Cal),
	}
	return FoodReq{Food: food}, nil
}

func DecodeGrpcUpdateFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateFoodRequest)
	food := Food{
		Name:    req.Name,
		Brand:   req.Brand,
		Amount:  float64(req.Amount),
		Unit:    req.Unit,
		Carb:    float64(req.Carb),
		Portein: float64(req.Portein),
		Fat:     float64(req.Fat),
		Cal:     float64(req.Cal),
	}
	return FoodReq{Food: food}, nil
}

func DecodeGrpcDeleteFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteFoodRequest)
	return GetFoodByIdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcDefaultResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(DefaultResp)

	return &pb.DefaultResponse{Result: resp.Result}, nil

}
