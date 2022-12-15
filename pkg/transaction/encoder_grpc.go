package transaction

import (
	"context"
	pb "go-api-grpc/pb/transaction"
)

type GRPCEncoder struct{}

func NewGRPCEncoder() GRPCEncoder {
	return GRPCEncoder{}
}

func (GRPCEncoder) Response(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseDTO)
	amount, _ := resp.Amount.Float64()
	return &pb.Response{
		ID:          resp.ID.String(),
		FromAccount: resp.FromAccount.String(),
		ToAccount:   resp.FromAccount.String(),
		Amount:      amount,
		Currency:    resp.Currency,
		CreatedAt:   resp.CreatedAt.String(),
	}, nil
}

func (e GRPCEncoder) ListResponse(context context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseListDTO)
	pbResp := pb.ResponseList{
		Page:         int64(resp.Page),
		PerPage:      int64(resp.PerPage),
		TotalPage:    int64(resp.TotalPages),
		TotalRecords: int64(resp.TotalRecords),
	}
	
	for _, item := range resp.Data {
		pbRespItem, err := e.Response(context, item)
		if err != nil {
			return nil, err
		}
		pbResp.Data = append(pbResp.Data, pbRespItem.(*pb.Response))
	}
	
	return pbResp, nil
}
