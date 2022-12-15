package boot

import (
	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go-api-grpc/configs"
	pbAccount "go-api-grpc/pb/account"
	pbTransaction "go-api-grpc/pb/transaction"
	"go-api-grpc/pkg/account"
	"go-api-grpc/pkg/transaction"
	"go-api-grpc/utils/database"
	"go-api-grpc/utils/validation"
	"google.golang.org/grpc"
	"net"
	"os"
)

func Grpc(cfg configs.GRPC, grpcListener net.Listener, logger kitLog.Logger) {
	// init configs
	cfg, err := configs.NewGRPCConfig()
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}
	
	// init DB
	db, err := database.NewPostgres(cfg.DB)
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}
	
	// validation service
	validator, err := validation.NewValidation()
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}
	
	// init mappers
	accountMapper := account.NewMapper()
	txMapper := transaction.NewMapper()
	
	// init repos
	accountRepo := account.NewRepository(db)
	txRepo := transaction.NewRepository(db)
	
	// init services
	accountService := account.NewService(accountRepo, accountMapper)
	txService := transaction.NewService(txRepo, txMapper)
	
	// init logging service
	accountLoggingService := account.NewLoggingService(logger, accountService)
	txLoggingService := transaction.NewLoggingService(logger, txService)
	
	// init endpoints
	accountEndpoint := account.NewEndpoint(validator, accountLoggingService)
	txEndpoint := transaction.NewEndpoint(validator, txLoggingService)
	
	// init grpc decoders
	accountDecoder := account.NewGRPCDecoder()
	txDecoder := transaction.NewGRPCDecoder()
	
	// init grpc encoders
	accountEncoder := account.NewGRPCEncoder()
	txEncoder := transaction.NewGRPCEncoder()
	
	// grpc servers
	accountServer := account.NewGRPCTransport(accountEndpoint, accountDecoder, accountEncoder)
	txServer := transaction.NewGRPCTransport(txEndpoint, txDecoder, txEncoder)
	
	baseServer := grpc.NewServer()
	pbAccount.RegisterAccountServer(baseServer, accountServer)
	pbTransaction.RegisterTransactionServer(baseServer, txServer)
	level.Info(logger).Log("msg", "Server started successfully 🚀")
	baseServer.Serve(grpcListener)
}