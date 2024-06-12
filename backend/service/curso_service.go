package services

import (
	cursoClient "Arqui_Soft_I/backend/clients/curso"
	materiaClient "Arqui_Soft_I/backend/clients/materia"
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
		materia, err := materiaClient.GetMateriaById(curso.Materia_id)
		if err != nil {
			return nil, e.NewBadRequestApiError("Failed to get materia for curso")
		}
		var cursoDto dto.CursoDto
		{
			cursoDto.ID = curso.ID
			cursoDto.Fecha_Inicio = curso.Fecha_Inicio
			cursoDto.Fecha_Fin = curso.Fecha_Fin
			cursoDto.Materia_id = curso.Materia_id
			cursoDto.Requisitos = curso.Requisitos
			cursoDto.Instructor = curso.Instructor
			cursoDto.Materia.ID = materia.ID
			cursoDto.Materia.Nombre = materia.Nombre
			cursoDto.Materia.Duracion = materia.Duracion
			cursoDto.Materia.Descripcion = materia.Descripcion
			cursoDto.Materia.Palabras_clave = materia.Palabras_clave
		}
		cursosDto = append(cursosDto, cursoDto)
	}

	return cursosDto, nil
}
