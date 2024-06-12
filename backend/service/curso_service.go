package services

import (
	cursoClient "Arqui_Soft_I/backend/clients/curso"
	"Arqui_Soft_I/backend/dto"
	"Arqui_Soft_I/backend/model"

	e "Arqui_Soft_I/backend/utils"
)

type cursoService struct{}

type cursoServiceInterface interface {
	GetCursos() (dto.Cursos, e.ApiError)
}

var (
	CursoService cursoServiceInterface
)

func init() {
	CursoService = &cursoService{}
}

func (s *cursoService) GetCursos() (dto.Cursos, e.ApiError) {

	var cursos model.Cursos = cursoClient.GetCursos()
	var cursosDto dto.Cursos

	for _, curso := range cursos {
		var cursoDto dto.CursoDto
		cursoDto.ID = curso.ID
		cursoDto.Fecha_Inicio = curso.Fecha_Inicio
		cursoDto.Fecha_Fin = curso.Fecha_Fin
		cursoDto.Materia_id = curso.Materia_id
		cursoDto.Requisitos = curso.Requisitos
		cursoDto.Instructor = curso.Instructor
		cursosDto = append(cursosDto, cursoDto)
	}

	return cursosDto, nil
}
