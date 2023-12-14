package repository

import (
	"muassisa-service/internal/models"
	"muassisa-service/internal/pkg/db"
	"muassisa-service/internal/pkg/logger"

	"go.uber.org/fx"
)

var NewRepository = fx.Provide(newRepository)

type IRepository interface {
	GetCourse(language int64) (transactions []*models.GetCourse, errs error)
	GeLanguage() (lang []models.Language, errs error)
	Photo(name string, photo []byte) error
	GetPhoto() (fileName string, errs error)
	GetPhotoName(id int64) (fileName []Photos, errs error)
	GetCourseNew() (transactions []models.GetCourseNew, errs error)
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

func newRepository(dp dependencies) IRepository {
	return &repository{
		Postgres: dp.Postgres,
		Logger:   dp.Logger,
	}
}
