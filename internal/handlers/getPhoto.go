package handlers

import (
	"encoding/json"
	"log"
	response "muassisa-service/internal/models"
	"net/http"
)

type Phot struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	photo string `json:"photo" gorm:"column:iamge"`
}

func (ch Handler) GetPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var language lang
		err := json.NewDecoder(r.Body).Decode(&language)
		if err != nil {
			ch.Logger.Error("request", response.SetError(err))
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		log.Println("language-------> ", language.Id)

		course, err := ch.svc.RepositoryInstance().GetPhoto(language.Id)
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		//var file Phot
		//for _, v := range course {
		//	log.Println("course------> ", v)
		//	file.Id = v.Id
		//	file.Name = v.Name
		//	// Открытие файла
		//	files, err := os.Open("assets/" + v.Photo)
		//	log.Println(file)
		//	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//	defer files.Close()
		//	file.Asd = files
		//	// Отправка изображения в ответ
		//	_, err = io.Copy(w, file.Asd)
		//	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//
		//}
		var pho Phot
		for _, v := range course {
			log.Println("course------>", v)
			pho.Id = v.Id
			pho.Name = v.Name
			pho.photo = "C:\\Users\\user\\Desktop\\muassisa-service\\assets\\" + v.Photo
		}
		log.Print("pho----->", pho)
		// Сериализация данных в JSON
		jsonData, err := json.Marshal(pho)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Установка заголовка ответа и отправка данных в ответе
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		//ch.Logger.Info("course: ", pho)
		//response.ToJson(w, http.StatusOK, map[string]interface{}{
		//	"response": pho,
		//})
	}
}
