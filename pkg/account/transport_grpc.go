package account

import (
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
	pb "go-wallet-api/pb/account/v1"
)

type GRPCTransport struct {
	create gt.Handler
	list   gt.Handler
	pb.UnimplementedAccountServiceServer
}

func NewGRPCTransport(endpoint Endpoint, decoder GRPCDecoder, encoder GRPCEncoder) pb.AccountServiceServer {
	return &GRPCTransport{
		create: gt.NewServer(
			endpoint.Create,
			decoder.CreateRequest,
			encoder.Response,
		),
		list: gt.NewServer(
			endpoint.List,
			decoder.ListRequest,
			encoder.ListResponse,
		),
	}
}

func (s *GRPCTransport) Create(ctx context.Context, req *pb.CreateRequest) (*pb.Response, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Response), nil
}

func (s *GRPCTransport) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListResponse), nil
}
