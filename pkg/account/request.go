package account

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateRequest struct {
	UserID   uuid.UUID       `json:"user_id" validate:"required"`
	Balance  decimal.Decimal `json:"balance" validate:"dge=0"`
	Currency string          `json:"currency" validate:"required,oneof=USD BTC ETC"`
}

type ListRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}
