package repository

func (r repository) Photo(name string, photo []byte) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `insert into photos (name, data) VALUES (?,?)`
	err := db.Exec(sqlQuery, name, photo)
	if err != nil {
		return err.Error
	}

	return nil
}
