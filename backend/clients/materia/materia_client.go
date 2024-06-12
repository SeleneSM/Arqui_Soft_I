package clients

import (
	"Arqui_Soft_I/backend/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetMateriaById(id int) (model.Materia, error) {
	var materia model.Materia
	result := Db.Where("id = ?", id).First(&materia)
	if result.Error != nil {
		// Loguear el error para debugging
		log.Error("Failed to find materia with ID:", id, "Error:", result.Error)
		return model.Materia{}, result.Error
	}

	log.Debug("Materia: ", materia)

	return materia, nil
}

func SearchMateria(keyword string) model.Materias {
	var materias model.Materias
	search_pattern := "%" + keyword + "%"

	Db.Where("nombre LIKE ? OR descripcion LIKE ? OR palabras_clave LIKE ?", search_pattern, search_pattern, search_pattern).Find(&materias)
	log.Debug("Materias disponibles: ", materias)

	return materias
}
