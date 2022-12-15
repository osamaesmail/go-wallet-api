package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	pb "go-api-grpc/pb/transaction/v1"
	"go-api-grpc/utils/pagination"
)

type GRPCDecoder struct{}

func NewGRPCDecoder() GRPCDecoder {
	return GRPCDecoder{}
}

func (GRPCDecoder) CreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateRequest)
	fromAccount, err := uuid.Parse(req.FromAccount)
	if err != nil {
		return nil, err
	}
	toAccount, err := uuid.Parse(req.ToAccount)
	if err != nil {
		return nil, err
	}
	amount := decimal.NewFromFloat(req.Amount)
	return CreateRequest{
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
	}, nil
}

func (GRPCDecoder) ListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ListRequest)
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	return ListRequest{
		UserID: userID,
		Request: pagination.Request{
			Page:    int(req.Page),
			PerPage: int(req.PerPage),
		},
	}, nil
}
