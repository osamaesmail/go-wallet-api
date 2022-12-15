package transaction

import (
	"github.com/go-kit/log"
	"time"
)

type LoggingService struct {
	logger log.Logger
	IService
}

func NewLoggingService(logger log.Logger, service IService) LoggingService {
	return LoggingService{logger, service}
}

func (s LoggingService) Create(req CreateRequest) (res ResponseDTO, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.IService.Create(req)
}

func (s LoggingService) List(req ListRequest) (res ResponseListDTO, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.IService.List(req)
}
