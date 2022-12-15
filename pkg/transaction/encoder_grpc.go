package transaction

import (
	"context"
	pb "go-wallet-api/pb/transaction/v1"
)

type GRPCEncoder struct{}

func NewGRPCEncoder() GRPCEncoder {
	return GRPCEncoder{}
}

func (GRPCEncoder) Response(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseDTO)
	amount, _ := resp.Amount.Float64()
	return &pb.Response{
		Id:          resp.ID.String(),
		FromAccount: resp.FromAccount.String(),
		ToAccount:   resp.FromAccount.String(),
		Amount:      amount,
		Currency:    resp.Currency,
		CreatedAt:   resp.CreatedAt.String(),
	}, nil
}

func (e GRPCEncoder) ListResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseListDTO)
	pbResp := &pb.ListResponse{
		Page:         int64(resp.Page),
		PerPage:      int64(resp.PerPage),
		TotalPage:    int64(resp.TotalPages),
		TotalRecords: int64(resp.TotalRecords),
	}
	
	for _, item := range resp.Data {
		pbRespItem, err := e.Response(ctx, item)
		if err != nil {
			return nil, err
		}
		pbResp.Data = append(pbResp.Data, pbRespItem.(*pb.Response))
	}
	
	return pbResp, nil
}
