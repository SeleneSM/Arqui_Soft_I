package services

import (
	comentarioClient "Arqui_Soft_I/backend/clients/comentario"
	"Arqui_Soft_I/backend/dto"
	"Arqui_Soft_I/backend/model"
	e "Arqui_Soft_I/backend/utils"

	"time"
)

type comentarioService struct{}

type comentarioServiceInterface interface {
	InsertComentario(comentarioDto dto.ComentarioDto) (dto.ComentarioDto, e.ApiError)
	GetComentariosPorCursoID(cursoID int) (model.Comentarios, e.ApiError)
}

var (
	ComentarioService comentarioServiceInterface
)

func init() {
	ComentarioService = &comentarioService{}
}

func (s *comentarioService) InsertComentario(comentarioDto dto.ComentarioDto) (dto.ComentarioDto, e.ApiError) {
	var comentario_DAO model.Comentario
	comentario_DAO.ID_curso = comentarioDto.CursoID
	comentario_DAO.ID_usuario = comentarioDto.UsuarioID
	comentario_DAO.Texto = comentarioDto.Texto
	comentario_DAO.Fecha_comentario = time.Now()
	comentario_DAO.Usuario.Username = comentarioDto.Usuario.Username

	comentario_DAO, err := comentarioClient.InsertComentario(comentario_DAO)
	if err != nil {
		return dto.ComentarioDto{}, e.NewInternalServerApiError("Error creando comentario", err)
	}
	comentarioDto.ID = comentario_DAO.ID

	return comentarioDto, nil
}

func (s *comentarioService) GetComentariosPorCursoID(cursoID int) (model.Comentarios, e.ApiError) {
	comentarios, err := comentarioClient.GetComentariosPorCursoID(cursoID)
	if err != nil {
		return nil, e.NewInternalServerApiError("Error obteniendo comentarios", err)
	}

	return comentarios, nil
}
