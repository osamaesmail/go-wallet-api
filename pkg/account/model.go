package account

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        uuid.UUID       `gorm:"primarykey"`
	UserID    uuid.UUID       `gorm:"uniqueIndex"`
	Balance   decimal.Decimal `gorm:"sql:decimal(20,8)"`
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// todo prevent many currencies in one account
// todo: add in DB level
func (a *Account) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
