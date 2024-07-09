package clients

import (
	"Arqui_Soft_I/backend/model" //Llama a model

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

func GetCursos() model.Cursos {
	var cursos model.Cursos
	Db.Find(&cursos)
	log.Debug("Cursos: ", cursos)
	return cursos
}

func InsertCurso(curso model.Curso) model.Curso {
	result := Db.Create(&curso)

	if result.Error != nil {

		//TODO Manage Errors

		log.Error("")
	}
	log.Debug("Curso creado: ", curso.ID)
	return curso
}
