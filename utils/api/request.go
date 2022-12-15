package api

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewJsonRequestDecoder[T any](T) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, request *http.Request) (interface{}, error) {
		var req T
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		return req, err
	}
}
