package repository

import (
	"log"
	"muassisa-service/internal/models"
)

func (r repository) GetLesson(id string) (lesson []models.GetLesson, errs error) {
	db := r.Postgres.GetPostgresConnection()

	sqlQuery := `select id,course_id, name, title, description, link, time, created_at from curriculum where course_id=?;`
	err := db.Raw(sqlQuery, id).Scan(&lesson)
	if err != nil {
		return lesson, err.Error
	}
	log.Println(lesson)
	return lesson, nil
}
