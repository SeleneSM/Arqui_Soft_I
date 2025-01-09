package services

import (
	inscripcionClient "Arqui_Soft_I/clients/inscripcion"
	"Arqui_Soft_I/dto"
	"Arqui_Soft_I/model"
	e "Arqui_Soft_I/utils"

	log "github.com/sirupsen/logrus"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	InscribirUsuario(inscrip dto.InscripcionDto) (dto.InscripcionDto, e.ApiError)
	GetInscripcionesByUser(userId int) (dto.Inscripciones, e.ApiError)
}

var (
	InscripcionService inscripcionServiceInterface
)

func init() {
	InscripcionService = &inscripcionService{}
}

// Con esta funcion convierto el DTO en DAO
func (s *inscripcionService) InscribirUsuario(inscrip dto.InscripcionDto) (dto.InscripcionDto, e.ApiError) {
	var inscripcion_DAO model.Inscripcion
	inscripcion_DAO.ID_curso = inscrip.ID_curso
	inscripcion_DAO.ID_usuario = inscrip.ID_usuario
	inscripcion_DAO.Fecha_Inscripcion = inscrip.Fecha_Inscripcion
	//Llamo a cliente para realizar la inscripcion
	log.Debug("user_id", inscripcion_DAO.ID_usuario)
	inscripcion_DAO = inscripcionClient.InscribirUsuario(inscripcion_DAO)

	inscrip.ID = inscripcion_DAO.ID

	return inscrip, nil

	//Finalmente, retorno la inscripcion del usuario al curso
}

func (s *inscripcionService) GetInscripcionesByUser(userId int) (dto.Inscripciones, e.ApiError) {

	var inscripciones model.Inscripciones = inscripcionClient.GetInscripcionesByUser(userId)
	var inscripcionesDto dto.Inscripciones

	for _, inscripcion := range inscripciones {
		var inscripcionDto dto.InscripcionDto

		inscripcionDto.Fecha_Inscripcion = inscripcion.Fecha_Inscripcion
		inscripcionDto.ID_curso = inscripcion.ID_curso
		inscripcionDto.ID_usuario = inscripcion.ID_usuario
		inscripcionDto.ID = inscripcion.ID
		inscripcionesDto = append(inscripcionesDto, inscripcionDto)

	}

	return inscripcionesDto, nil
}
