package repository

func (r repository) AddPhoto(name, url string) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := "insert into photo (name, iamge) VALUES (?,?);"
	err := db.Exec(sqlQuery, name, url)
	if err != nil {
		return err.Error
	}

	return nil
}
