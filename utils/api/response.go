package api

import (
	"context"
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	resp := &ErrorResponse{Errors: []Error{}}

	switch cErr := err.(type) {
	case Error:
		resp.Errors = []Error{cErr}
	case validator.ValidationErrors:
		errs := handleValidationErrors(cErr)
		resp.Errors = errs
	default:
		resp.Errors = []Error{
			{
				Message: err.Error(),
				Code:    strconv.Itoa(http.StatusInternalServerError),
			},
		}
	}

	err2 := kithttp.EncodeJSONResponse(ctx, w, resp)
	if err2 != nil {
		kithttp.DefaultErrorEncoder(ctx, err, w)
	}
}

func handleValidationErrors(errors validator.ValidationErrors) []Error {
	var errorResponse = make([]Error, 0, len(errors))

	for _, err := range errors {
		e := validationErrorToText(err)
		errorResponse = append(errorResponse, e)
	}

	return errorResponse
}

// todo, other validation formats
func validationErrorToText(e validator.FieldError) Error {
	var detail string
	switch e.Tag() {
	case "required":
		detail = fmt.Sprintf("%s is required", toSnakeCase(e.Field()))
	default:
		detail = fmt.Sprintf("%s is not valid", toSnakeCase(e.Field()))
	}
	return Error{
		Code:    strconv.Itoa(http.StatusUnprocessableEntity),
		Source:  toSnakeCase(e.Field()),
		Message: detail,
	}
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
