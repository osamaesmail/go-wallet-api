package account

import (
	"context"
	pb "go-api-grpc/pb/account/v1"
)

type GRPCEncoder struct{}

func NewGRPCEncoder() GRPCEncoder {
	return GRPCEncoder{}
}

func (GRPCEncoder) Response(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseDTO)
	balance, _ := resp.Balance.Float64()
	return &pb.Response{
		Id:       resp.ID.String(),
		UserId:   resp.UserID.String(),
		Balance:  balance,
		Currency: resp.Currency,
	}, nil
}

func (e GRPCEncoder) ListResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(ResponseListDTO)
	pbResp := &pb.ListResponse{}

	for _, item := range resp {
		pbRespItem, err := e.Response(ctx, item)
		if err != nil {
			return nil, err
		}
		pbResp.Data = append(pbResp.Data, pbRespItem.(*pb.Response))
	}

	return pbResp, nil
}
