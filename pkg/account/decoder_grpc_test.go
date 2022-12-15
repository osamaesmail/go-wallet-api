package account_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	pb "go-wallet-api/pb/account/v1"
	"go-wallet-api/pkg/account"
	"testing"
)

func setupDecoder() account.GRPCDecoder {
	return account.NewGRPCDecoder()
}

func TestDecoderGRPC_CreateRequest(t *testing.T) {
	decoder := setupDecoder()

	// test data
	grpcReq := &pb.CreateRequest{
		UserId:   "ef907719-4820-4917-b161-4cd3f418c6aa",
		Balance:  200,
		Currency: "USD",
	}
	userID, _ := uuid.Parse(grpcReq.UserId)
	req := account.CreateRequest{
		UserID:   userID,
		Balance:  decimal.NewFromFloat(grpcReq.Balance),
		Currency: grpcReq.Currency,
	}

	resp, err := decoder.CreateRequest(context.Background(), grpcReq)
	assert.NoError(t, err)
	assert.Equal(t, req, resp)
}
