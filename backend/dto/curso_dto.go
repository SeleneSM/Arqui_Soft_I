package dto

import (
	"time"
)

type CursoDto struct {
	ID           int       `json:"id"`
	Fecha_Inicio time.Time `json:"fecha_inicio"`
	Fecha_Fin    time.Time `json:"fecha_fin"`
	Materia_id   int       `json:"materia_id"`
	Requisitos   string    `json:"requisitos"`
	Instructor   string    `json:"instructor"`
}

type Cursos []CursoDto
