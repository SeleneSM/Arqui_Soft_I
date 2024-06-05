package services

import (
	materiaClient "Arqui_Soft_I/backend/clients/materia"
	"Arqui_Soft_I/backend/dto"
	"Arqui_Soft_I/backend/model"
	e "Arqui_Soft_I/backend/utils"
)

type materiaService struct{}
type materiaServiceInterface interface {
	SearchMateria(palabras_clave string) (dto.Materias, e.ApiError)
}

var (
	MateriasService materiaServiceInterface
)

func init() {
	MateriasService = &materiaService{}
}

func (s *materiaService) SearchMateria(palabras_clave string) (dto.Materias, e.ApiError) {

	var materias model.Materias = materiaClient.SearchMateria(palabras_clave)
	var materiasDto dto.Materias

	for _, materia := range materias {
		var materiaDto dto.MateriaDto
		materiaDto.Nombre = materia.Nombre
		materiaDto.ID = materia.ID
		materiaDto.Duracion = materia.Duracion
		materiaDto.Descripcion = materia.Descripcion
		materiaDto.Palabras_clave = materia.Palabras_clave
		materiasDto = append(materiasDto, materiaDto)
	}

	return materiasDto, nil
}

func (s *materiaService) GetMateriaById(id int) (dto.MateriaDto, e.ApiError) {

	var materia model.Materia = materiaClient.GetMateriaById(id)
	var materiaDto dto.MateriaDto

	if materia.ID == 0 {
		return materiaDto, e.NewBadRequestApiError("materia not found")
	}

	materiaDto.ID = materia.ID
	materiaDto.Nombre = materia.Nombre
	materiaDto.Duracion = materia.Duracion
	materiaDto.Descripcion = materia.Descripcion
	materiaDto.Palabras_clave = materia.Palabras_clave

	return materiaDto, nil
}
