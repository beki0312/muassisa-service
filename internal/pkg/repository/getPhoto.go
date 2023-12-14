package repository

import (
	"fmt"
	"log"
	"os"
)

type Photos struct {
	ID    int64
	Name  string `gorm:"column:name"`
	Image []byte `gorm:"column:data"`
}

func (r repository) GetPhoto() (fileName string, errs error) {
	db := r.Postgres.GetPostgresConnection()
	err := db.AutoMigrate(Photos{})
	if err != nil {
		log.Fatal(err)
	}

	// Выборка фото
	//var photos []Photo

	var photos []Photos
	err = db.Table("photos").Find(&photos).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("photos---> ", photos)
	// Итерирование по фото и сохранение в файлы
	for _, photo := range photos {
		fileName, err = savePhotoToFile(photo.Image, fmt.Sprintf("photo_%d.jpg", photo.ID))
		if err != nil {
			log.Fatal(err)
		}
	}

	return fileName, nil
}

// Функция сохранения фото в файл
func savePhotoToFile(photo []byte, fileName string) (files string, err error) {
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
