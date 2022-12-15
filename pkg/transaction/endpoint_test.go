package transaction_test

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	transactionMock "go-api-grpc/mocks/pkg/transaction"
	"go-api-grpc/pkg/transaction"
	"go-api-grpc/utils/pagination"
	"go-api-grpc/utils/validation"
	"testing"
)

type endpointMocks struct {
	Service *transactionMock.IService
}

func setupEndpoint() (transaction.Endpoint, endpointMocks) {
	valid, _ := validation.NewValidation()
	service := new(transactionMock.IService)
	endpoint := transaction.NewEndpoint(valid, service)
	mocks := endpointMocks{
		Service: service,
	}
	return endpoint, mocks
}

func TestEndpoint_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			// setup
			endpoint, m := setupEndpoint()

			// test data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			respDTO := transaction.ResponseDTO{
				ID:          uuid.New(),
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}

			// mocks
			m.Service.On("Create", req).Return(respDTO, nil)

			// call method
			resp, err := endpoint.Create(context.Background(), req)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, respDTO, resp)
			m.Service.AssertExpectations(t)
		},
	)

	t.Run(
		"error - validation", func(t *testing.T) {
			// setup
			endpoint, _ := setupEndpoint()

			// test data
			req := transaction.CreateRequest{
				FromAccount: uuid.Nil,
				ToAccount:   uuid.Nil,
				Amount:      decimal.New(-100, 0),
			}

			// call method
			_, err := endpoint.Create(context.Background(), req)

			// assert
			assert.NotNil(t, err)
			assert.IsType(t, validator.ValidationErrors{}, err)
			assert.Len(t, err.(validator.ValidationErrors), 3)
		},
	)

	t.Run(
		"error - service", func(t *testing.T) {
			// setup
			endpoint, m := setupEndpoint()

			// test data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}

			// mocks
			m.Service.On("Create", req).Return(transaction.ResponseDTO{}, errors.New("err"))

			// call method
			_, err := endpoint.Create(context.Background(), req)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, "err", err.Error())
			m.Service.AssertExpectations(t)
		},
	)
}

func TestEndpoint_List(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			// setup
			endpoint, m := setupEndpoint()

			// test data
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}
			respDTO := transaction.ResponseDTO{
				ID:          uuid.New(),
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
				Currency:    "USD",
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

			// mocks
			m.Service.On("List", req).Return(respListDTO, nil)

			// call method
			resp, err := endpoint.List(context.Background(), req)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, respListDTO, resp)
			m.Service.AssertExpectations(t)
		},
	)

	t.Run(
		"error - validation", func(t *testing.T) {
			// setup
			endpoint, _ := setupEndpoint()

			// test data
			req := transaction.ListRequest{
				UserID: uuid.Nil,
			}

			// call method
			_, err := endpoint.List(context.Background(), req)

			// assert
			assert.NotNil(t, err)
			assert.IsType(t, validator.ValidationErrors{}, err)
			assert.Len(t, err.(validator.ValidationErrors), 1)
		},
	)

	t.Run(
		"error - service", func(t *testing.T) {
			// setup
			endpoint, m := setupEndpoint()

			// test data
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}

			// mocks
			m.Service.On("List", req).Return(transaction.ResponseListDTO{}, errors.New("err"))

			// call method
			_, err := endpoint.List(context.Background(), req)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, "err", err.Error())
			m.Service.AssertExpectations(t)
		},
	)
}
