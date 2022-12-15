package account

import (
	"go-wallet-api/utils/api"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(model Account) (resp Account, err error)
	List(req ListRequest) (resp []Account, err error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(model Account) (Account, error) {
	err := r.db.Create(&model).Error
	// todo handle duplicate error
	return model, err
}

func (r Repository) List(req ListRequest) (resp []Account, err error) {
	err = r.db.Where("user_id", req.UserID).Find(&resp).Error
	if len(resp) == 0 {
		return resp, api.Error{Code: "404", Message: "User not found"}
	}
	return
}
