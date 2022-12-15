package configs

import (
	"go-wallet-api/utils/config"
)

type Rest struct {
	Port int64 `mapstructure:"REST_PORT" validate:"required"`
	DB   DB    `mapstructure:",squash"`
}

func NewRestConfig() (api Rest, err error) {
	err = config.DefaultEnv().Bind(&api)
	return
}
