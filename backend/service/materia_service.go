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

/*
func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

		var hotels model.Hotels = hotelClient.GetHotels()
		var hotelsDto dto.HotelsDto

		for _, hotel := range hotels {
			var hotelDto dto.HotelDto
			hotelDto.Name = hotel.Nombre
			hotelDto.CantHabitaciones = hotel.CantHab
			hotelDto.Id = hotel.ID
			hotelDto.Desc = hotel.Descripcion
			hotelsDto = append(hotelsDto, hotelDto)
		}

		return hotelsDto, nil
	}
*/
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
