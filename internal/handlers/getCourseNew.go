package handlers

import (
	"encoding/json"
	"muassisa-service/internal/models"
	"net/http"
)

func (ch Handler) GetCourseNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		course, err := ch.svc.RepositoryInstance().GetCourseNew()
		if err != nil {
			http.Error(w, "Ошибка", http.StatusInternalServerError)
			return
		}
		var p []models.GetCourseNew
		for _, v := range course {
			var b models.GetCourseNew
			b.Id = v.Id
			b.CourseName = v.CourseName
			b.Title = v.Title
			b.CourseDescription = v.CourseDescription
			b.Amount = v.Amount
			b.ImageName = v.ImageName
			b.Image = "http://schooltaj.na4u.ru/assets/" + v.Image
			b.Category = v.Category
			b.Status = v.Status
			b.InstructorName = v.InstructorName
			b.InsrtuctorUnvon = v.InsrtuctorUnvon
			b.InstructorDescription = v.InstructorDescription
			b.Lectures = v.Lectures
			b.Quizzes = v.Quizzes
			b.Duration = v.Duration
			b.SkillLevel = v.SkillLevel
			b.Language = v.Language
			b.Students = v.Students
			b.Assessments = v.Assessments
			b.DateBegin = v.DateBegin
			b.DateEnd = v.DateEnd
			b.CreatedAt = v.CreatedAt
			p = append(p, b)
		}

		// Преобразуем структуру в JSON
		jsonResponse, err := json.Marshal(p)
		if err != nil {
			http.Error(w, "Ошибка при преобразовании в JSON", http.StatusInternalServerError)
			return
		}
		// Устанавливаем заголовок Content-Type на application/json
		w.Header().Set("Content-Type", "application/json")
		// Отправляем JSON-ответ
		w.Write(jsonResponse)
	}
}
