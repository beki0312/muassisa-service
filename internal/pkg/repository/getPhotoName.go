package repository

import (
	"log"
)

func (r repository) GetPhotoName(id int64) (fileName []Photos, errs error) {
	db := r.Postgres.GetPostgresConnection()
	query := `select id, name,data from photos where id = ?`
	//var person []Photos
	err := db.Raw(query, id).Scan(&fileName)

	log.Println("person------>", fileName)
	if err != nil {
		log.Println(err)
		return
	}
	return fileName, nil
}
