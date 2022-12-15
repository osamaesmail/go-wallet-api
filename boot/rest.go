package boot

import (
	"github.com/gin-gonic/gin"
	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"go-api-grpc/configs"
	"go-api-grpc/pkg/account"
	"go-api-grpc/pkg/transaction"
	"go-api-grpc/utils/api"
	"go-api-grpc/utils/database"
	"go-api-grpc/utils/validation"
	"log"
	"os"
)

func Rest(cfg configs.Rest) *gin.Engine {
	// init router
	engine := gin.Default()
	apiV1 := engine.Group("/api/v1")
	
	// validation service
	validator, err := validation.NewValidation()
	if err != nil {
		log.Fatal(err)
	}
	
	// kit logger init
	var logger kitLog.Logger
	logger = kitLog.NewLogfmtLogger(os.Stderr)
	logger = kitLog.With(logger, "listen", "8081", "caller", kitLog.DefaultCaller)
	
	// kit options for http transport
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(api.EncodeError),
	}
	
	// init DB
	db, err := database.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatal(err)
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
	
	// init controllers
	accountControllerHttp := account.NewHTTPTransport(opts, accountEndpoint)
	txControllerHttp := transaction.NewHTTPTransport(opts, txEndpoint)
	
	// init routes
	account.NewHttpRouter(apiV1, accountControllerHttp)
	transaction.NewHttpRouter(apiV1, txControllerHttp)
	
	return engine
}
