package dto

import (
	"time"
)

type CursoDto struct {
	ID           int       `json:"id"`
	Fecha_Inicio time.Time `json:"fecha_inicio"`
	Fecha_Fin    time.Time `json:"fecha_fin"`
	MateriaID    int       `json:"materia"`
}

type Cursos []CursoDto
