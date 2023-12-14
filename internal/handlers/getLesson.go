package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (ch Handler) GetLessn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)[`id`]
		lesson, err := ch.svc.RepositoryInstance().GetLesson(id)
		if err != nil {
			http.Error(w, "Ошибка при преобразовании в JSON", http.StatusInternalServerError)
			return
		}

		log.Println(lesson)
		// Преобразуем структуру в JSON
		jsonResponse, err := json.Marshal(lesson)
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
