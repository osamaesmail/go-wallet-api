package transaction

import (
	"go-api-grpc/utils/slice"
)

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m Mapper) CreateRequestToModel(req CreateRequest) Transaction {
	return Transaction{
		FromAccount: req.FromAccount,
		ToAccount:   req.ToAccount,
		Amount:      req.Amount,
	}
}

func (m Mapper) ModelDTOToDTO(t DTO) ResponseDTO {
	return ResponseDTO(t)
}

func (m Mapper) ModelToDTO(t Transaction) ResponseDTO {
	return ResponseDTO{
		ID:          t.ID,
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Amount:      t.Amount,
		CreatedAt:   t.CreatedAt,
	}
}

func (m Mapper) ModelsDTOToListDTO(list []DTO) ResponseListDTO {
	resp := ResponseListDTO{}
	resp.Data = slice.Map(
		list, func(i int, t DTO) ResponseDTO {
			return m.ModelDTOToDTO(t)
		},
	)
	return resp
}

func (m Mapper) ModelsToListDTO(list []Transaction) ResponseListDTO {
	resp := ResponseListDTO{}
	resp.Data = slice.Map(
		list, func(i int, t Transaction) ResponseDTO {
			return m.ModelToDTO(t)
		},
	)
	return resp
}
