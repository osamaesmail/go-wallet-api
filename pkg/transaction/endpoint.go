package transaction

import (
	"context"
	"github.com/go-playground/validator/v10"
)

type IEndpoint interface {
	Create(ctx context.Context, request interface{}) (interface{}, error)
	List(ctx context.Context, request interface{}) (interface{}, error)
}

type Endpoint struct {
	validator *validator.Validate
	service   IService
}

func NewEndpoint(validate *validator.Validate, service IService) Endpoint {
	return Endpoint{validate, service}
}

func (e Endpoint) Create(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(CreateRequest)
	err := e.validator.Struct(req)
	if err != nil {
		return nil, err
	}
	res, err := e.service.Create(req)
	return res, err
}

func (e Endpoint) List(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(ListRequest)
	err := e.validator.Struct(req)
	if err != nil {
		return nil, err
	}
	res, err := e.service.List(req)
	return res, err
}
