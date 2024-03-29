package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-wallet-api/configs"
	"go-wallet-api/utils/database"
	"go-wallet-api/utils/migrations"
)

func Migrations() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting migrations")
			runMigrations(args)
		},
	}
}

func runMigrations(args []string) {
	dbConfig, err := configs.NewDBConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewPostgres(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	migrations.RunWithArgs(db, args)
}
