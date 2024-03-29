package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	pb "go-wallet-api/pb/account/v1"
)

type GRPCDecoder struct{}

func NewGRPCDecoder() GRPCDecoder {
	return GRPCDecoder{}
}

func (GRPCDecoder) CreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateRequest)
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	balance := decimal.NewFromFloat(req.Balance)
	return CreateRequest{
		UserID:   userID,
		Balance:  balance,
		Currency: req.Currency,
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
	}, nil
}
