package repository

func (r repository) AddedCourse(name, title, description, imageName, image string, amount, language int64) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := "insert into course(name, title, description, amount, image_name, image, language) values (?,?,?,?,?,?,?);"
	err := db.Exec(sqlQuery, name, title, description, amount, imageName, image, language)
	if err != nil {
		return err.Error
	}

	return nil
}
