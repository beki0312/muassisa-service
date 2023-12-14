package handlers

import (
	"log"
	response "muassisa-service/internal/models"
	"net/http"
	"os"
)

func (ch Handler) GetPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//params := mux.Vars(r)
		//photoID := params["id"]
		//id, _ := strconv.Atoi(photoID)
		photoData, err := ch.svc.RepositoryInstance().GetPhoto()
		log.Println("photoData----> ", photoData)
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}

		// Установка заголовка Content-Type
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename="+photoData)
		w.Header().Set("Content-Transfer-Encoding", "binary")
		// Отправка фото в ответе
		//_, err = w.Write(photoData)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//ch.Logger.Info("course: ", photoData)
		//response.ToJson(w, http.StatusOK, map[string]interface{}{
		//	"response": photoData,
		//})

		f, _ := os.ReadFile(photoData)
		defer os.Remove(photoData)
		defer w.Write(f)
		w.WriteHeader(200)
	}
}
