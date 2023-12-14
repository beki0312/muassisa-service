package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	response "muassisa-service/internal/models"
	"net/http"
)

func (ch Handler) Photo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // Максимальный размер загружаемого файла (10 МБ)

		// Получение имени фото из формы
		name := r.FormValue("name")

		// Чтение данных фото из переданного файла
		file, _, err := r.FormFile("photo")
		if err != nil {
			log.Println(err)
			http.Error(w, "Ошибка загрузки фото", http.StatusBadRequest)
			return
		}
		defer file.Close()

		photoData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			http.Error(w, "Ошибка чтения фото", http.StatusBadRequest)
			return
		}

		// Сохранение данных фото в базу данных или другое необходимое действие
		err = ch.svc.RepositoryInstance().Photo(name, photoData)
		if err != nil {
			ch.Logger.Error("response", err)
			log.Println(err)
			http.Error(w, "Ошибка сохранения фото", http.StatusInternalServerError)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}

		ch.Logger.Info("course: ", err)

		// Возвращаем успешный ответ
		response := map[string]string{"message": "Фото успешно сохранено"}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	}
}
