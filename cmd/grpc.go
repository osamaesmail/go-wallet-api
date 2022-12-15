package cmd

import (
	"fmt"
	kitLog "github.com/go-kit/kit/log"
	
	"github.com/go-kit/kit/log/level"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-api-grpc/boot"
	"go-api-grpc/configs"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func GRPC() *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "Run gRPC Server",
		Run:   runGrpc,
	}
}

func runGrpc(cmd *cobra.Command, args []string) {
	// init configs
	cfg, err := configs.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}
	
	// logger
	var logger kitLog.Logger
	logger = kitLog.NewJSONLogger(os.Stdout)
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)
	logger = kitLog.With(logger, "caller", kitLog.DefaultCaller)
	
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}
	
	go boot.Grpc(cfg, grpcListener, logger)
	
	level.Error(logger).Log("exit", <-errs)
}
