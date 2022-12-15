package transaction

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go-wallet-api/utils/pagination"
	"time"
)

type ResponseDTO struct {
	ID          uuid.UUID `gorm:"primarykey"`
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	Amount      decimal.Decimal
	Currency    string
	CreatedAt   time.Time
}

type ResponseListDTO struct {
	Data []ResponseDTO `json:"data"`
	pagination.Response
}
