package configs

import (
	"go-api-grpc/utils/config"
)

type DB struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     int64  `mapstructure:"DB_PORT" validate:"required,number"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	User     string `mapstructure:"DB_USER" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
}

func NewDBConfig() (db DB, err error) {
	err = config.DefaultEnv().Bind(&db)
	return
}
