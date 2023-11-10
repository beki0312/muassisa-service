package repository

import (
	"muassisa-service/internal/models"
	"muassisa-service/internal/pkg/db"
	"muassisa-service/internal/pkg/logger"

	"go.uber.org/fx"
)

var NewRepositoryTransaction = fx.Provide(newRepositoryTransaction)

type ITransactionRepository interface {
	AddTransactionsPos(transactions []*models.AddTransaction) (errResponse models.ErrorResponse)
	GetCourse(language int64) (transactions []*models.GetCourse, errs error)
	GeLanguage() (lang []models.Language, errs error)
	UpdateStatus(transId string) (errResponse models.ErrorResponse)
}

type dependencies struct {
	fx.In
	Postgres db.IPostgres
	Logger   logger.ILogger
}

type repository struct {
	Postgres db.IPostgres
	Logger   logger.ILogger
}

func newRepositoryTransaction(dp dependencies) ITransactionRepository {
	return &repository{
		Postgres: dp.Postgres,
		Logger:   dp.Logger,
	}
}
