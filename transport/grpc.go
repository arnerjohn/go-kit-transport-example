package transport

import (
	"context"
	"github.com/arnerjohn/transport-example/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	uppercase grpctransport.Handler
	count     grpctransport.Handler
}

func (s *grpcServer) Uppercase(ctx context.Context, req *pb.Request) (*pb.UppercaseResponse, error) {
	_, rep, err := s.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*pb.UppercaseResponse), nil
}

func (s *grpcServer) Count(ctx context.Context, req *pb.Request) (*pb.CountResponse, error) {
	_, rep, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*pb.CountResponse), nil
}

func NewGRPCServer(endpoints EndpointSet) pb.TransportExampleServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{
		uppercase: grpctransport.NewServer(
			endpoints.UppercaseEndpoint,
			decodeGRPCUppercaseRequest,
			encodeGRPCUppercaseResponse,
			append(options)...,
		),
		count: grpctransport.NewServer(
			endpoints.CountEndpoint,
			decodeGRPCCountRequest,
			encodeGRPCCountResponse,
			append(options)...,
		),
	}
}

func decodeGRPCUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Request)
	return UppercaseRequest{Input: req.Input}, nil
}

func encodeGRPCUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UppercaseResponse)
	return &pb.UppercaseResponse{Output: resp.Output, Err: resp.Err}, nil
}

func decodeGRPCCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Request)
	return CountRequest{Input: req.Input}, nil
}

func encodeGRPCCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CountResponse)
	return &pb.CountResponse{Output: int64(resp.Output)}, nil
}
