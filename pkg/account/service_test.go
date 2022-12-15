package account_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	accountMock "go-wallet-api/mocks/pkg/account"
	"go-wallet-api/pkg/account"
	"testing"
)

type serviceMocks struct {
	Repo *accountMock.IRepository
}

func setupService() (account.Service, serviceMocks) {
	mapper := account.NewMapper()
	repo := new(accountMock.IRepository)
	service := account.NewService(repo, mapper)
	mocks := serviceMocks{
		Repo: repo,
	}
	return service, mocks
}

func TestService_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			// setup
			service, m := setupService()

			// test data
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			model := account.Account{
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}
			newModel := model
			newModel.ID = uuid.New()
			respDTO := account.ResponseDTO{
				ID:       newModel.ID,
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}

			m.Repo.On("Create", model).Return(newModel, nil)

			// call method
			resp, err := service.Create(req)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, respDTO, resp)
			m.Repo.AssertExpectations(t)
		},
	)

	t.Run(
		"error", func(t *testing.T) {
			// setup
			service, m := setupService()

			// test data
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			model := account.Account{
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}

			// mocks
			m.Repo.On("Create", model).Return(account.Account{}, errors.New("err"))

			// call method
			_, err := service.Create(req)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "err")
			m.Repo.AssertExpectations(t)
		},
	)
}

func TestService_List(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			// setup
			service, m := setupService()

			// test data
			req := account.ListRequest{
				UserID: uuid.New(),
			}
			model := account.Account{
				ID:       req.UserID,
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			modesls := []account.Account{model, model}
			respDTO := account.ResponseDTO{
				ID:       model.ID,
				UserID:   model.UserID,
				Balance:  model.Balance,
				Currency: model.Currency,
			}
			respListDTO := account.ResponseListDTO{respDTO, respDTO}

			m.Repo.On("List", req).Return(modesls, nil)

			// call method
			resp, err := service.List(req)

			// assert
			assert.Nil(t, err)
			assert.Len(t, resp, 2)
			assert.Equal(t, respListDTO, resp)
			m.Repo.AssertExpectations(t)
		},
	)

	t.Run(
		"error", func(t *testing.T) {
			// setup
			service, m := setupService()

			// test data
			req := account.ListRequest{
				UserID: uuid.New(),
			}

			// mocks
			m.Repo.On("List", req).Return([]account.Account{}, errors.New("err"))

			// call method
			_, err := service.List(req)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "err")
			m.Repo.AssertExpectations(t)
		},
	)
}
