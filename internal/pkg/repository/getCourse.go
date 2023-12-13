package repository

import (
	"log"
	"muassisa-service/internal/models"
)

func (r repository) GetCourse(language int64) (transactions []*models.GetCourse, errs error) {
	db := r.Postgres.GetPostgresConnection()
	log.Println(language)

	sqlQuery := `select id, name, title, description,amount,image_name,image, date_begin, date_end, created_at from course c where active and language=?`
	err := db.Raw(sqlQuery, language).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
