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
	getbodys   gt.Handler
	createbody gt.Handler
	updatebody gt.Handler
	deletebody gt.Handler
}

func NewGRPCServer(endpoints GrpcEndpoints, logger log.Logger) pb.BodyServiceServer {
	return &grpcBodyServer{
		getbyid: gt.NewServer(
			endpoints.GetById,
			DecodeGrpcGetBodyByIdRequest,
			EncodeGrpcResponse,
		),
		getbodys: gt.NewServer(
			endpoints.GetBodys,
			DecodeGrpcGetBodysequest,
			EncodeGrpcGetBodysResponse,
		),
		createbody: gt.NewServer(
			endpoints.CreateBody,
			DecodeGrpcCreateBodyRequest,
			EncodeGrpcDefaultResponse,
		),
		updatebody: gt.NewServer(
			endpoints.UpdateBody,
			DecodeGrpcUpdateBodyRequest,
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

func (s *grpcBodyServer) GetBodys(ctx context.Context, req *pb.GetBodysRequest) (*pb.GetBodysResponse, error) {
	_, resp, err := s.getbodys.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetBodysResponse), nil
}

func (s *grpcBodyServer) CreateBody(ctx context.Context, req *pb.CreateBodyRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.createbody.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.DefaultResponse), nil
}

func (s *grpcBodyServer) UpdateBody(ctx context.Context, req *pb.UpdateBodyRequest) (*pb.DefaultResponse, error) {
	_, resp, err := s.updatebody.ServeGRPC(ctx, req)
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

	result := pb.Body{
		Bodyid:  body.Id,
		Weight:  float32(body.Weight),
		Muscle:  float32(body.Muscle),
		FatRate: float32(body.FatRate),
	}

	return &pb.GetBodyByIdResponse{
		Body: &result,
	}, nil

}

func DecodeGrpcGetBodysequest(_ context.Context, request interface{}) (interface{}, error) {
	// request.(*pb.GetBodysRequest)
	return GetBodysReq{}, nil
}

func EncodeGrpcGetBodysResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetBodysGrpcResp)

	if resp.Body == nil {
		return nil, errors.New("id not found")
	}

	// var bodys []Body
	bodylist := resp.Body.([]interface{})

	var results []*pb.Body
	for _, item := range bodylist {
		body := item.(Body)
		result := pb.Body{
			Bodyid:  body.Id,
			Weight:  float32(body.Weight),
			Muscle:  float32(body.Muscle),
			FatRate: float32(body.FatRate),
		}
		results = append(results, &result)
	}

	return &pb.GetBodysResponse{
		Body: results,
	}, nil
}

func DecodeGrpcCreateBodyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateBodyRequest)

	body := Body{
		Weight:  float64(req.Body.Weight),
		Muscle:  float64(req.Body.Muscle),
		FatRate: float64(req.Body.FatRate),
	}
	return BodyReq{Body: body}, nil
}

func DecodeGrpcUpdateBodyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateBodyRequest)

	body := Body{
		Id:      req.Body.Bodyid,
		Weight:  float64(req.Body.Weight),
		Muscle:  float64(req.Body.Muscle),
		FatRate: float64(req.Body.FatRate),
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
