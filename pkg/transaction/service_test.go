package transaction_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	transactionMock "go-wallet-api/mocks/pkg/transaction"
	"go-wallet-api/pkg/transaction"
	"go-wallet-api/utils/pagination"
	"testing"
)

type serviceMocks struct {
	Repo *transactionMock.IRepository
}

func setupService() (transaction.Service, serviceMocks) {
	mapper := transaction.NewMapper()
	repo := new(transactionMock.IRepository)
	service := transaction.NewService(repo, mapper)
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
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			respDTO := transaction.ResponseDTO{
				ID:          newModel.ID,
				FromAccount: newModel.FromAccount,
				ToAccount:   newModel.ToAccount,
				Amount:      newModel.Amount,
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
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			
			// mocks
			m.Repo.On("Create", model).Return(transaction.Transaction{}, errors.New("err"))
			
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
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}
			modelDto := transaction.DTO{
				ID:          uuid.New(),
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			modelsDto := []transaction.DTO{modelDto, modelDto}
			respDTO := transaction.ResponseDTO{
				ID:          modelDto.ID,
				FromAccount: modelDto.FromAccount,
				ToAccount:   modelDto.ToAccount,
				Amount:      modelDto.Amount,
			}
			respListDTO := transaction.ResponseListDTO{
				Data: []transaction.ResponseDTO{respDTO, respDTO},
				Response: pagination.Response{
					Page:         1,
					PerPage:      10,
					TotalPages:   1,
					TotalRecords: 2,
				},
			}
			
			m.Repo.On("List", req).Return(modelsDto, int64(2), nil)
			
			// call method
			resp, err := service.List(req)
			
			// assert
			assert.Nil(t, err)
			assert.Equal(t, respListDTO, resp)
			m.Repo.AssertExpectations(t)
		},
	)
	
	t.Run(
		"error", func(t *testing.T) {
			// setup
			service, m := setupService()
			
			// test data
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}
			
			// mocks
			m.Repo.On("List", req).Return([]transaction.DTO{}, int64(0), errors.New("err"))
			
			// call method
			_, err := service.List(req)
			
			// assert
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "err")
			m.Repo.AssertExpectations(t)
		},
	)
}
