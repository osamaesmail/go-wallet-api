package transaction

import (
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
	pb "go-api-grpc/pb/transaction/v1"
)

type GRPCTransport struct {
	create gt.Handler
	list   gt.Handler
	pb.UnimplementedTransactionServiceServer
}

func NewGRPCTransport(endpoint Endpoint, decoder GRPCDecoder, encoder GRPCEncoder) pb.TransactionServiceServer {
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
