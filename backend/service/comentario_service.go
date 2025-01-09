package services

import (
	comentarioClient "Arqui_Soft_I/clients/comentario"
	userClient "Arqui_Soft_I/clients/user"

	"Arqui_Soft_I/dto"
	"Arqui_Soft_I/model"
	e "Arqui_Soft_I/utils"

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
	usuar, err := userClient.GetUserById(comentarioDto.UsuarioID)
	if err != nil {
		return dto.ComentarioDto{}, e.NewInternalServerApiError("Error geting user by id", err)
	}
	//DAO que se carga en la base de datos
	comentario_DAO.ID_curso = comentarioDto.CursoID
	comentario_DAO.ID_usuario = comentarioDto.UsuarioID
	comentario_DAO.Texto = comentarioDto.Texto
	comentario_DAO.Fecha_comentario = time.Now()
	comentario_DAO.Usuario.Username = usuar.Username
	comentario_DAO.Usuario.Apellido = usuar.Apellido
	comentario_DAO.Usuario.Rol = usuar.Rol
	comentario_DAO.Usuario.Nombre = usuar.Nombre

	//dto que devuelve la funcion
	comentarioDto.Fecha_comentario = time.Now()
	comentarioDto.Usuario.Username = usuar.Username
	comentarioDto.Usuario.Apellido = usuar.Apellido
	comentarioDto.Usuario.Password = usuar.Password
	comentarioDto.Usuario.Rol = usuar.Rol
	comentarioDto.Usuario.ID = usuar.ID
	comentarioDto.Usuario.Nombre = usuar.Nombre

	comentario_DAO, err = comentarioClient.InsertComentario(comentario_DAO)
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
