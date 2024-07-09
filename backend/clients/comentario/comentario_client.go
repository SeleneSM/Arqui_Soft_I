package client

import (
	"Arqui_Soft_I/backend/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertComentario(comentario model.Comentario) (model.Comentario, error) {
	if err := Db.Create(&comentario).Error; err != nil {
		log.Println("Error creando comentario:", err)
		return model.Comentario{}, err
	}
	return comentario, nil
}

func GetComentariosPorCursoID(cursoID int) (model.Comentarios, error) {
	var comentarios model.Comentarios
	if err := Db.Where("curso_id = ?", cursoID).Find(&comentarios).Error; err != nil {
		log.Println("Error obteniendo comentarios para curso ID:", cursoID, err)
		return nil, err
	}
	return comentarios, nil
}
