package configs

import (
	"go-api-grpc/utils/config"
)

type Rest struct {
	Port int64 `mapstructure:"REST_PORT" validate:"required"`
	DB   DB    `mapstructure:",squash"`
}

func NewRestConfig() (api Rest, err error) {
	err = config.DefaultEnv().Bind(&api)
	return
}
