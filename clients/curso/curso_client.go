package clients

import (
	"Arqui_Soft_I/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCursoById(id int) model.Curso {
	var curso model.Curso

	Db.Where("id = ?", id).First(&curso)
	log.Debug("Curso: ", curso)

	return curso
}
