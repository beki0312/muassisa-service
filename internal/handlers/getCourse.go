package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"muassisa-service/internal/models"
	"net/http"
	"strconv"
)

type lang struct {
	Id int64 `json:"id"`
}

func (ch Handler) GetCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)[`id`]
		num, _ := strconv.Atoi(id)

		course, err := ch.svc.RepositoryInstance().GetCourse(int64(num))
		if err != nil {
			http.Error(w, "Ошибка", http.StatusInternalServerError)
			return
		}
		var p []models.GetCourse
		for _, v := range course {
			var b models.GetCourse
			b.Id = v.Id
			b.Name = v.Name
			b.Title = v.Title
			b.Description = v.Description
			b.Amount = v.Amount
			b.ImageName = v.ImageName
			b.Image = "http://schooltaj.na4u.ru/assets/" + v.Image
			b.DateBegin = v.DateBegin
			b.DateEnd = v.DateEnd
			b.Created_At = v.Created_At
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
