package account

import "go-api-grpc/utils/slice"

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m Mapper) CreateRequestToModel(req CreateRequest) Account {
	return Account{
		UserID:   req.UserID,
		Balance:  req.Balance,
		Currency: req.Currency,
	}
}

func (m Mapper) ModelToDTO(t Account) ResponseDTO {
	return ResponseDTO{
		ID:       t.ID,
		UserID:   t.UserID,
		Balance:  t.Balance,
		Currency: t.Currency,
	}
}

func (m Mapper) ModelsToListDTO(list []Account) []ResponseDTO {
	return slice.Map(
		list, func(i int, t Account) ResponseDTO {
			return m.ModelToDTO(t)
		},
	)
}
