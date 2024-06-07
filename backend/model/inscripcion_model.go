package model

import (
	"time"
)

type Inscripcion struct {
	ID                int       `gorm:"primaryKey"`
	ID_curso          int       `gorm:"foreignKey"`
	ID_usuario        int       `gorm:"foreignKey"`
	Fecha_Inscripcion time.Time `gorm:"type:date"`
}

type Inscripciones []Inscripcion
