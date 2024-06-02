package model

import (
	"time"
)

type Curso struct {
	ID           int       `gorm:"primaryKey"`
	Fecha_Inicio time.Time `gorm:"type:date"`
	Fecha_Fin    time.Time `gorm:"type:date"`
	Materia_id   int       `gorm:"type:int"`
}

type Cursos []Curso