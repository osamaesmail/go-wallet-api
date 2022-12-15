package transaction

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Amount      decimal.Decimal
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	CreatedAt   time.Time
}

type DTO struct {
	ID          uuid.UUID `gorm:"primarykey"`
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	Amount      decimal.Decimal
	Currency    string
	CreatedAt   time.Time
}

// todo: add in DB level
func (a *Transaction) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
