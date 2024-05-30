package clients

import (
	"Arqui_Soft_I/backend/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetMateriaById(id int) model.Materia {
	var materia model.Materia

	Db.Where("id = ?", id).First(&materia)
	log.Debug("Materia: ", materia)

	return materia
}
