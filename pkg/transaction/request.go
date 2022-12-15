package transaction

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go-wallet-api/utils/pagination"
)

type CreateRequest struct {
	FromAccount uuid.UUID       `json:"from_account" validate:"required"`
	ToAccount   uuid.UUID       `json:"to_account" validate:"required"`
	Amount      decimal.Decimal `json:"amount" validate:"dge=0.00000001"`
}

type ListRequest struct {
	pagination.Request
	UserID uuid.UUID `json:"user_id" validate:"required"`
}
