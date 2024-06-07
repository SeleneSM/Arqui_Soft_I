package dto

import (
	"time"
)

type InscripcionDto struct {
	ID                int       `json:"id"` //id inscripcion
	ID_curso          int       `json:"id_curso"`
	ID_usuario        int       `json:"id_usuario"`
	Fecha_Inscripcion time.Time `json:"fecha_inscripcion"`
}

type Inscripciones []InscripcionDto
