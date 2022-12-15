package account_test

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	accountMock "go-api-grpc/mocks/pkg/account"
	"go-api-grpc/pkg/account"
	"go-api-grpc/utils/validation"
	"testing"
)

type endpointMocks struct {
	Service *accountMock.IService
}

func setupEndpoint() (account.Endpoint, endpointMocks) {
	valid, _ := validation.NewValidation()
	service := new(accountMock.IService)
	endpoint := account.NewEndpoint(valid, service)
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
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			respDTO := account.ResponseDTO{
				ID:       uuid.New(),
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}
			
			// mocks
			m.Service.On("Create", req).Return(respDTO, nil)
			
			// call method
			resp, err := endpoint.Create(nil, req)
			
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
			req := account.CreateRequest{
				UserID:   uuid.Nil,
				Balance:  decimal.New(-100, 0),
				Currency: "Foo",
			}
			
			// call method
			_, err := endpoint.Create(nil, req)
			
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
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			
			// mocks
			m.Service.On("Create", req).Return(account.ResponseDTO{}, errors.New("err"))
			
			// call method
			_, err := endpoint.Create(nil, req)
			
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
			req := account.ListRequest{
				UserID: uuid.New(),
			}
			respDTO := account.ResponseDTO{
				ID:       uuid.New(),
				UserID:   req.UserID,
				Balance:  decimal.New(100, 0),
				Currency: "USD",
			}
			respListDTO := account.ResponseListDTO{respDTO, respDTO}
			
			// mocks
			m.Service.On("List", req).Return(respListDTO, nil)
			
			// call method
			resp, err := endpoint.List(nil, req)
			
			// assert
			assert.Nil(t, err)
			assert.Len(t, resp, 2)
			assert.Equal(t, respListDTO, resp)
			m.Service.AssertExpectations(t)
		},
	)
	
	t.Run(
		"error - validation", func(t *testing.T) {
			// setup
			endpoint, _ := setupEndpoint()
			
			// test data
			req := account.ListRequest{
				UserID: uuid.Nil,
			}
			
			// call method
			_, err := endpoint.List(nil, req)
			
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
			req := account.ListRequest{
				UserID: uuid.New(),
			}
			
			// mocks
			m.Service.On("List", req).Return(account.ResponseListDTO{}, errors.New("err"))
			
			// call method
			_, err := endpoint.List(nil, req)
			
			// assert
			assert.NotNil(t, err)
			assert.Equal(t, "err", err.Error())
			m.Service.AssertExpectations(t)
		},
	)
}
