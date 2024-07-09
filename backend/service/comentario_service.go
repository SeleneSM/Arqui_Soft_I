package services

import (
	comentarioClient "Arqui_Soft_I/backend/clients/comentario"
	"Arqui_Soft_I/backend/dto"
	"Arqui_Soft_I/backend/model"
	e "Arqui_Soft_I/backend/utils"

	log "github.com/sirupsen/logrus"
	"time"
)
type comentarioService struct{}

type comentarioServiceInterface interface {
	InsertComentario(comentarioDto dto.ComentarioDto) (model.Comentario, e.ApiError)
	
}

var (
	comentarioService comentarioServiceInterface
)

func init() {
	comentarioService = &comentarioService{}
}


func InsertComentario(comentarioDto dto.ComentarioDto) (model.Comentario, e.ApiError) {
    var comentario_DAO = model.Comentario
	comentario_DAO.ID_curso =  comentarioDto.CursoID,
	comentario_DAO.ID_usuario = comentarioDto.UsuarioID,
	comentario_DAO.Texto =    comentarioDto.Texto,
	comentario_DAO.Fecha_comentario =     time.Now(),
	comentario_DAO.Usuario.Username= comentarioDto.Usuario.Username ,


    createdComentario, err := client.CrearComentario(comentario)
    if err != nil {
        return model.Comentario{}, e.NewInternalServerApiError("Error creando comentario", err)
    }

    return createdComentario, nil
}