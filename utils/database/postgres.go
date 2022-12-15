package database

import (
	"fmt"
	"go-api-grpc/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres(config configs.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)
	return gorm.Open(
		postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
}
