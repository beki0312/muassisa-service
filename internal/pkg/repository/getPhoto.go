package repository

import (
	"log"
	"muassisa-service/internal/models"
)

func (r repository) GetPhoto(id int64) (transactions []*models.Photoo, errs error) {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := "select * from photo where id=?"
	err := db.Raw(sqlQuery, id).Scan(&transactions).Error
	for _, v := range transactions {
		log.Println("transactions--> ", v)
	}
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
