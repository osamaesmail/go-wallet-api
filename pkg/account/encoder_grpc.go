package account

import (
	"context"
	pb "go-api-grpc/pb/account"
)

type GRPCEncoder struct{}

func NewGRPCEncoder() GRPCEncoder {
	return GRPCEncoder{}
}

func (GRPCEncoder) Response(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseDTO)
	balance, _ := resp.Balance.Float64()
	return &pb.Response{
		ID:       resp.ID.String(),
		UserID:   resp.UserID.String(),
		Balance:  balance,
		Currency: resp.Currency,
	}, nil
}

func (e GRPCEncoder) ListResponse(context context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseListDTO)
	pbResp := pb.ResponseList{}
	
	for _, item := range resp {
		pbRespItem, err := e.Response(context, item)
		if err != nil {
			return nil, err
		}
		pbResp.Data = append(pbResp.Data, pbRespItem.(*pb.Response))
	}
	
	return pbResp, nil
}