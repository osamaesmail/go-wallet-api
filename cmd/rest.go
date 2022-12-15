package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go-wallet-api/boot"
	"go-wallet-api/configs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Rest() *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run API server",
		Run:   runRest,
	}
}

func runRest(cmd *cobra.Command, args []string) {
	// init configs
	cfg, err := configs.NewRestConfig()
	if err != nil {
		log.Fatal(err)
	}
	
	app := boot.Rest(cfg)
	
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: app,
	}
	
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
	<-quit
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
