package services

import (
	cursoClient "Arqui_Soft_I/clients/curso"
	//materiaClient "Arqui_Soft_I/backend/clients/materia"
	"Arqui_Soft_I/dto"
	"Arqui_Soft_I/model"
	e "Arqui_Soft_I/utils"
)

type cursoService struct{}

type cursoServiceInterface interface {
	GetCursos() (dto.Cursos, e.ApiError)
	GetCursoById(id int) (dto.CursoDto, e.ApiError)
	InsertCurso(cursoDto dto.CursoDto) (dto.CursoDto, e.ApiError)
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
		/*materia, err := materiaClient.GetMateriaById(curso.Materia_id)
		if err != nil {
			return nil, e.NewBadRequestApiError("Failed to get materia for curso")
		}*/
		var cursoDto dto.CursoDto
		{
			cursoDto.ID = curso.ID
			cursoDto.Fecha_Inicio = curso.Fecha_Inicio
			cursoDto.Fecha_Fin = curso.Fecha_Fin
			cursoDto.Materia_id = curso.Materia_id
			cursoDto.Requisitos = curso.Requisitos
			cursoDto.Instructor = curso.Instructor
			//cursoDto.Materia.ID = materia.ID
			//cursoDto.Materia.Nombre = materia.Nombre
			//cursoDto.Materia.Duracion = materia.Duracion
			//cursoDto.Materia.Descripcion = materia.Descripcion
			//cursoDto.Materia.Palabras_clave = materia.Palabras_clave
		}
		cursosDto = append(cursosDto, cursoDto)
	}

	return cursosDto, nil
}

func (s *cursoService) GetCursoById(id int) (dto.CursoDto, e.ApiError) {

	var curso model.Curso = cursoClient.GetCursoById(id)
	var cursoDto dto.CursoDto

	if curso.ID == 0 {
		return cursoDto, e.NewBadRequestApiError("curse not found")
	}

	cursoDto.ID = curso.ID
	cursoDto.Fecha_Inicio = curso.Fecha_Inicio
	cursoDto.Fecha_Fin = curso.Fecha_Fin
	cursoDto.Materia_id = curso.Materia_id
	//cursoDto.Materia.Descripcion = curso.Materia.Descripcion
	//cursoDto.Materia.Nombre = curso.Materia.Nombre
	//cursoDto.Materia.Palabras_clave = curso.Materia.Palabras_clave
	//cursoDto.Materia.Duracion = curso.Materia.Duracion
	//cursoDto.Materia.ID = curso.Materia.ID
	cursoDto.Requisitos = curso.Requisitos
	cursoDto.Instructor = curso.Instructor

	return cursoDto, nil
}

func (s *cursoService) InsertCurso(cursoDto dto.CursoDto) (dto.CursoDto, e.ApiError) {

	var curso model.Curso

	curso.Fecha_Inicio = cursoDto.Fecha_Inicio
	curso.Fecha_Fin = cursoDto.Fecha_Fin
	curso.Materia_id = cursoDto.Materia_id
	//curso.Materia.Descripcion = cursoDto.Materia.Descripcion
	//curso.Materia.Nombre = cursoDto.Materia.Nombre
	//curso.Materia.Palabras_clave = cursoDto.Materia.Palabras_clave
	//curso.Materia.Duracion = cursoDto.Materia.Duracion
	//curso.Materia.ID = cursoDto.Materia.ID
	curso.Requisitos = cursoDto.Requisitos
	curso.Instructor = cursoDto.Instructor

	curso = cursoClient.InsertCurso(curso)

	cursoDto.ID = curso.ID

	return cursoDto, nil
}
