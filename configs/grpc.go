package configs

import (
	"go-api-grpc/utils/config"
)

type GRPC struct {
	Port int64 `mapstructure:"GRPC_PORT" validate:"required"`
	DB   DB    `mapstructure:",squash"`
}

func NewGRPCConfig() (api GRPC, err error) {
	err = config.DefaultEnv().Bind(&api)
	return
}
