package repository

import (
	"log"
	"muassisa-service/internal/models"
)

func (r repository) GetCourseNew() (transactions []models.GetCourseNew, errs error) {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `select c.id,c.name,c.title,c.description,c.amount,c.image_name,c.image,c.category,c.status,i.name,i.unvon,i.description,f.lectures,f.quizzes,f.duration,f.skill_level,f.language,f.students,f.assessments,c.date_begin,c.date_end 
from course c
inner join instructor i on c.course_id=i.id 
inner join course_features f on c.id=f.course_id`
	err := db.Raw(sqlQuery).Scan(&transactions).Error
	log.Println("trans---- ", transactions)
	if err != nil {
		log.Println("errrr  ", err)
		return transactions, err
	}

	return transactions, nil
}
