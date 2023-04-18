package service

import (
	"context"
	"errors"
	pb "go-kit-demo/body/proto"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcBodyServer struct {
	getbyid    gt.Handler
	createbody gt.Handler
	deletebody gt.Handler
}

func NewGRPCServer(endpoints GrpcEndpoints, logger log.Logger) pb.BodyServiceServer {
	return &grpcBodyServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGrpcGetBodyByIdRequest,
			EncodeGrpcResponse,
		),
		createbody: gt.NewServer(
			endpoints.CreateBody,
			DecodeGrpcCreateBodyRequest,
			EncodeGrpcDefaultResponse,
		),
		deletebody: gt.NewServer(
			endpoints.DeleteBody,
			DecodeGrpcDeleteBodyRequest,
			EncodeGrpcDefaultResponse,
		),
	}
}

func (s *grpcBodyServer) GetBodyById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetBodyByIdResponse, error) {
	_, resp, err := s.getbyid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetBodyByIdResponse), nil
}

func (s *grpcBodyServer) CreateBody(ctx context.Context, req *pb.CreateBodyRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.createbody.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcBodyServer) DeleteBody(ctx context.Context, req *pb.GetByIdRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.deletebody.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func DecodeGrpcGetBodyByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetByIdRequest)
	return IdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetBodyByIdGrpcResp)

	if resp.Body == nil {
		return nil, errors.New("id not found")
	}

	body := resp.Body.(Body)
	return &pb.GetBodyByIdResponse{
		Bodyid:  body.Id,
		Weight:  float32(body.Weight),
		Muscle:  float32(body.Muscle),
		FatRate: float32(body.FatRate),
	}, nil

}

func DecodeGrpcCreateBodyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateBodyRequest)

	body := Body{
		Weight:  float64(req.Weight),
		Muscle:  float64(req.Muscle),
		FatRate: float64(req.FatRate),
	}
	return BodyReq{Body: body}, nil
}

func DecodeGrpcDeleteBodyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetByIdRequest)
	return IdGrpcReq{Id: req.Id}, nil
}

func EncodeGrpcDefaultResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(DefaultResp)

	return &pb.DefaultResponse{Result: resp.Result}, nil

}
