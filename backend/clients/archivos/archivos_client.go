package client

import (
	Model "Arqui_Soft_I/model"
	e "Arqui_Soft_I/utils"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func CreateFile(file Model.Archivo) (Model.Archivo, e.ApiError) {
	result := Db.Create(&file)

	if result.Error != nil {
		log.Error("Error al crear el file")
		log.Error(result.Error)
		return file, e.NewBadRequestApiError("Error al crear el file")
	}

	return file, nil
}

func SaveFile(fileName string, fileContent []byte) (string, e.ApiError) {
	uploadDir := "./uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	filePath := uploadDir + "/" + fileName
	err := os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		log.Error("Error al guardar el file ")
		log.Error(err)
		return "", e.NewInternalServerApiError("Error al guardar el file", err)
	}

	return filePath, nil
}
