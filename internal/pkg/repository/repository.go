package repository

import (
	"go.uber.org/fx"
	"muassisa-service/internal/models"
	"muassisa-service/internal/pkg/db"
)

var NewRepository = fx.Provide(newRepository)

type IRepository interface {
	GetCourse(language int64) (transactions []*models.GetCourse, errs error)
	GeLanguage() (lang []models.Language, errs error)
	AddedCourse(name, title, description, imageName, image string, amount, language int64) error
	GetCourseNew() (transactions []models.GetCourseNew, errs error)
	GetLesson(id string) (transactions []models.GetLesson, errs error)
}

type dependencies struct {
	fx.In
	Postgres db.IPostgres
	//Logger   logger.ILogger
}

type repository struct {
	Postgres db.IPostgres
	//Logger   logger.ILogger
}

func newRepository(dp dependencies) IRepository {
	return &repository{
		Postgres: dp.Postgres,
		//Logger:   dp.Logger,
	}
}
