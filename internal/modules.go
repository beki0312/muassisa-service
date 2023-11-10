package internal

import (
	"go.uber.org/fx"
	"muassisa-service/internal/config"
	"muassisa-service/internal/handlers"
	"muassisa-service/internal/pkg/db"
	"muassisa-service/internal/pkg/logger"
	"muassisa-service/internal/pkg/repository"
	"muassisa-service/internal/pkg/service"
)

var Modules = fx.Options(
	service.NewService,
	config.NewConfig,
	logger.NewLogger,
	handlers.NewTransactionHandler,
	repository.NewRepositoryTransaction,
	db.NewPostgres,
)
