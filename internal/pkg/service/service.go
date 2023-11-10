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
	RepositoryInstance() repository.IRepository
}

type dependencies struct {
	fx.In
	Config     config.IConfig
	Repository repository.IRepository
	Logger     logger.ILogger
}

type service struct {
	Config     config.IConfig
	Repository repository.IRepository
	Logger     logger.ILogger
}

func newService(d dependencies) IService {
	return &service{
		d.Config,
		d.Repository,
		d.Logger,
	}
}

func (s service) ConfigInstance() config.IConfig {
	return s.Config
}

func (s service) LoggerInstance() logger.ILogger {
	return s.Logger
}

func (s service) RepositoryInstance() repository.IRepository {
	return s.Repository
}
