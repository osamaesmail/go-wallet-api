package transaction

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"go-api-grpc/utils/api"
	"net/http"
)

type HTTPTransport struct {
	options  []kithttp.ServerOption
	endpoint IEndpoint
}

func NewHTTPTransport(
	options []kithttp.ServerOption,
	endpoint IEndpoint,
) HTTPTransport {
	return HTTPTransport{options, endpoint}
}

func (t HTTPTransport) Create() http.Handler {
	return kithttp.NewServer(
		t.endpoint.Create,
		api.NewJSONRequestDecoder(CreateRequest{}),
		kithttp.EncodeJSONResponse,
		t.options...,
	)
}

func (t HTTPTransport) List() http.Handler {
	return kithttp.NewServer(
		t.endpoint.List,
		api.NewJSONRequestDecoder(ListRequest{}),
		kithttp.EncodeJSONResponse,
		t.options...,
	)
}
