package service

import (
	"go.uber.org/fx"
	"muassisa-service/internal/config"
	"muassisa-service/internal/pkg/logger"
	"muassisa-service/internal/pkg/repository"
)

var NewService = fx.Provide(newService)

type IService interface {
	ConfigInstance() config.IConfig
	LoggerInstance() logger.ILogger
	TransactionRepositoryInstance() repository.ITransactionRepository
}

type dependencies struct {
	fx.In
	Config                config.IConfig
	TransactionRepository repository.ITransactionRepository
	Logger                logger.ILogger
}

type service struct {
	Config                config.IConfig
	TransactionRepository repository.ITransactionRepository
	Logger                logger.ILogger
}

func newService(d dependencies) IService {
	return &service{
		d.Config,
		d.TransactionRepository,
		d.Logger,
	}
}

func (s service) ConfigInstance() config.IConfig {
	return s.Config
}

func (s service) LoggerInstance() logger.ILogger {
	return s.Logger
}

func (s service) TransactionRepositoryInstance() repository.ITransactionRepository {
	return s.TransactionRepository
}
