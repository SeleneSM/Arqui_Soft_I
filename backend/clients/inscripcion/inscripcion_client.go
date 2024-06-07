package client

import (
	"Arqui_Soft_I/backend/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InscribirUsuario(inscripcion model.Inscripcion) model.Inscripcion {
	result := Db.Create(&inscripcion)
	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Incripcion Creada: ", inscripcion.ID)

	return inscripcion

}

func GetInscripcionesByUser(userId int) model.Inscripciones {
	var inscripciones model.Inscripciones
	Db.Where("id_usuario = ?", userId).Find(&inscripciones)
	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}
