package service

import (
	"context"
	"errors"
	endpoint "go-kit-demo/food/endpoint/grpc"
	pb "go-kit-demo/food/proto"
	svc "go-kit-demo/food/service"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcFoodServer struct {
	getbyid    gt.Handler
	getfoods   gt.Handler
	createfood gt.Handler
	updatefood gt.Handler
	deletefood gt.Handler
}

func NewGRPCServer(endpoints endpoint.GrpcEndpoints, logger log.Logger) pb.FoodServiceServer {
	return &grpcFoodServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGrpcGetFoodByIdRequest,
			EncodeGrpcResponse,
		),
		getfoods: gt.NewServer(
			endpoints.GetFoods,
			DecodeGrpcGetFoodsequest,
			EncodeGrpcGetFoodsResponse,
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

func (s *grpcFoodServer) GetFoods(ctx context.Context, req *pb.GetFoodsRequest) (*pb.GetFoodsResponse, error) {
	_, resp, err := s.getfoods.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetFoodsResponse), nil
}

func DecodeGrpcGetFoodsequest(_ context.Context, request interface{}) (interface{}, error) {
	// request.(*pb.GetFoodsRequest)
	return endpoint.GetFoodsReq{}, nil
}

func EncodeGrpcGetFoodsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetFoodsGrpcResp)

	if resp.Food == nil {
		return nil, errors.New("id not found")
	}

	foodlist := resp.Food.([]interface{})

	var results []*pb.Food
	for _, item := range foodlist {
		food := item.(svc.Food)
		result := pb.Food{
			Foodid:  food.Foodid,
			Name:    food.Name,
			Brand:   food.Brand,
			Amount:  float32(food.Amount),
			Unit:    food.Unit,
			Carb:    float32(food.Carb),
			Portein: float32(food.Portein),
			Fat:     float32(food.Fat),
			Cal:     float32(food.Cal),
		}
		results = append(results, &result)
	}

	return &pb.GetFoodsResponse{
		Food: results,
	}, nil
}

func (s *grpcFoodServer) CreateFood(ctx context.Context, req *pb.FoodRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.createfood.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcFoodServer) UpdateFood(ctx context.Context, req *pb.FoodRequest) (*pb.DefaultResponse, error) {
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
	return endpoint.GetFoodByIdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetFoodByIdGrpcResp)

	if resp.Food == nil {
		return nil, errors.New("id not found")
	}

	food := resp.Food.(svc.Food)

	result := pb.Food{
		Foodid:  food.Foodid,
		Name:    food.Name,
		Brand:   food.Brand,
		Amount:  float32(food.Amount),
		Unit:    food.Unit,
		Carb:    float32(food.Carb),
		Portein: float32(food.Portein),
		Fat:     float32(food.Fat),
		Cal:     float32(food.Cal),
	}

	return &pb.GetFoodByIdResponse{
		Food: &result,
	}, nil

}

func DecodeGrpcCreateFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FoodRequest)

	food := svc.Food{
		Name:    req.Food.Name,
		Brand:   req.Food.Brand,
		Amount:  float64(req.Food.Amount),
		Unit:    req.Food.Unit,
		Carb:    float64(req.Food.Carb),
		Portein: float64(req.Food.Portein),
		Fat:     float64(req.Food.Fat),
		Cal:     float64(req.Food.Cal),
	}

	return endpoint.FoodReq{Food: food}, nil
}

func DecodeGrpcUpdateFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FoodRequest)
	food := svc.Food{
		Foodid:  req.Food.Foodid,
		Name:    req.Food.Name,
		Brand:   req.Food.Brand,
		Amount:  float64(req.Food.Amount),
		Unit:    req.Food.Unit,
		Carb:    float64(req.Food.Carb),
		Portein: float64(req.Food.Portein),
		Fat:     float64(req.Food.Fat),
		Cal:     float64(req.Food.Cal),
	}
	return endpoint.FoodReq{Food: food}, nil
}

func DecodeGrpcDeleteFoodRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteFoodRequest)
	return endpoint.GetFoodByIdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcDefaultResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DefaultResp)

	return &pb.DefaultResponse{Result: resp.Result}, nil

}
