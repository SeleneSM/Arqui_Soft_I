package client

import (
	"Arqui_Soft_I/backend/model" //Llama a model
	"errors"

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InscribirUsuario(inscripcion model.Inscripcion) model.Inscripcion {

	var existingInscripcion model.Inscripcion
	if err := Db.Where("ID_curso = ? AND ID_usuario = ?", inscripcion.ID_curso, inscripcion.ID_usuario).First(&existingInscripcion).Error; err == nil {
		log.Info("Inscripción ya hecha anteriormente: Curso ID:", inscripcion.ID_curso, "Usuario ID:", inscripcion.ID_usuario)
		return model.Inscripcion{}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Failed to query inscription:", err.Error())
		return model.Inscripcion{}
	}
	result := Db.Create(&inscripcion)
	if result.Error != nil {
		log.Error("Failed to create inscription:", result.Error)
		return model.Inscripcion{}
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
