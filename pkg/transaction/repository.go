package transaction

import (
	"go-wallet-api/pkg/account"
	"go-wallet-api/utils/api"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(model Transaction) (resp Transaction, err error)
	List(req ListRequest) (resp []DTO, total int64, err error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

// todo prevent different currency transaction
func (r Repository) Create(model Transaction) (resp Transaction, err error) {
	tx := r.db.Begin()
	var fromAccount account.Account

	err = tx.Create(&model).Error
	if err != nil {
		tx.Rollback()
		return
	}
	subtractRes := tx.Model(account.Account{ID: model.FromAccount}).
		Where("balance - ? >= 0", model.Amount).
		Update("balance", gorm.Expr("balance - ?", model.Amount))
	if subtractRes.Error != nil {
		tx.Rollback()
		err = subtractRes.Error
		return
	}
	if subtractRes.RowsAffected != 1 {
		tx.Rollback()
		err = api.Error{Code: "400", Message: "From account not found or insufficient funds"}
		return
	}
	err = subtractRes.Find(&fromAccount).Error
	if err != nil {
		tx.Rollback()
		return
	}
	addRes := tx.Model(account.Account{ID: model.ToAccount}).
		Where("currency = ?", fromAccount.Currency).
		Where("id != ?", model.FromAccount).
		Update("balance", gorm.Expr("balance + ?", model.Amount))
	if addRes.Error != nil {
		tx.Rollback()
		err = addRes.Error
		return
	}
	if addRes.RowsAffected != 1 {
		tx.Rollback()
		err = api.Error{Code: "400", Message: "To Account not found or currency not matched"}
		return
	}
	err = tx.Commit().Error
	return model, err
}

func (r Repository) List(req ListRequest) (models []DTO, total int64, err error) {
	q := r.db.Model(&Transaction{}).
		Joins("JOIN accounts a_from ON transactions.from_account = a_from.id").
		Joins("JOIN accounts a_to ON transactions.to_account = a_to.id").
		Where("a_from.user_id", req.UserID).
		Or("a_to.user_id", req.UserID)

	err = q.Count(&total).Error
	if err != nil {
		return
	}

	err = q.Select("transactions.*", "a_from.currency currency").
		Offset(req.GetOffset()).Limit(req.GetPerPage()).Find(&models).Error

	return
}
