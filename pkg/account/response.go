package account

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ResponseDTO struct {
	ID       uuid.UUID       `json:"id"`
	UserID   uuid.UUID       `json:"user_id"`
	Balance  decimal.Decimal `json:"balance"`
	Currency string          `json:"currency"`
}

type ResponseListDTO []ResponseDTO
