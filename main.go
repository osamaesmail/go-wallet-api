package main

import (
	"github.com/spf13/cobra"
	"go-wallet-api/cmd"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "app",
		Short: "Run App Commands",
	}
	
	rootCmd.AddCommand(
		cmd.Rest(),
		cmd.GRPC(),
		cmd.Migrations(),
	)
	
	cobra.CheckErr(rootCmd.Execute())
}
