package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	response "muassisa-service/internal/models"
	"net/http"
	"os"
	"strconv"
)

type Phots struct {
	ID    int64
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (ch Handler) GetPhotoName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		photoID := params["id"]
		id, _ := strconv.Atoi(photoID)
		photoData, err := ch.svc.RepositoryInstance().GetPhotoName(int64(id))
		log.Println("photoData--11111--> ", photoData)
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		var fileName string
		for _, photo := range photoData {
			fileName, err = PhotoToFile(photo.Image, fmt.Sprintf("photo_%d.jpg", photo.ID))
			if err != nil {
				log.Fatal(err)
			}
		}
		var ph Phots
		for _, v := range photoData {
			ph.ID = v.ID
			ph.Name = v.Name
			ph.Image = fileName
		}
		//ch.Logger.Info("course: ", ph)
		//response.ToJson(w, http.StatusOK, map[string]interface{}{
		//	"response": ph,
		//})

		//w.Header().Set("Content-Type", "image/jpeg")
		//w.Header().Set("Content-Disposition", "attachment; filename="+ph.Image)
		//w.Header().Set("Content-Transfer-Encoding", "binary")
		//f, _ := os.ReadFile(ph.Image)
		//defer os.Remove(ph.Image)
		//defer w.Write(f)
		//w.WriteHeader(200)
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"response": ph,
		})
	}
}

// Функция сохранения фото в файл
func PhotoToFile(photo []byte, fileName string) (files string, err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(photo)
	if err != nil {
		return "", err
	}

	return fileName, err
}
