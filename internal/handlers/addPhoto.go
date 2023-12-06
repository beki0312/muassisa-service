package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	response "muassisa-service/internal/models"
	"net/http"
	"os"
	"path/filepath"
)

func (ch Handler) AddPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // Максимальный размер загружаемого файла (10 МБ)
		// Получение имени фото из формы
		name := r.FormValue("name")
		log.Println("name----> ", name)
		// Проверяем, получили ли файл изображения
		file, handler, err := r.FormFile("photo")
		if err != nil {
			http.Error(w, "Failed to receive image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Генерируем UUID
		uuid := uuid.New().String()

		// Получаем расширение файла
		ext := filepath.Ext(handler.Filename)
		// Формируем путь сохранения с UUID и расширением
		filePath := fmt.Sprintf("assets/%s%s", uuid, ext)
		// Создаем файл с путем сохранения
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		// Копируем содержимое файла изображения в новый файл
		_, err = io.Copy(f, file)
		if err != nil {
			log.Fatal(err)
		}
		err = ch.svc.RepositoryInstance().AddPhoto(name, uuid+ext)
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		ch.Logger.Info("course: ", "course")
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"response": "course",
		})
	}
}