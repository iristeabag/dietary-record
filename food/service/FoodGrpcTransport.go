package service

import (
	"context"
	"errors"
	pb "go-kit-demo/proto/food"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcFoodServer struct {
	getbyid gt.Handler
}

func NewGRPCServer(endpoints GrpcEndpoints, logger log.Logger) pb.FoodServiceServer {
	return &grpcFoodServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGetFoodByIdGrpcRequest,
			EncodeGrpcResponse,
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

func DecodeGetFoodByIdGrpcRequest(_ context.Context, request interface{}) (interface{}, error) {
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
		Amount:  food.Amount,
		Unit:    food.Unit,
		Carb:    food.Carb,
		Portein: food.Portein,
		Fat:     food.Fat,
		Cal:     food.Cal,
	}, nil

}
