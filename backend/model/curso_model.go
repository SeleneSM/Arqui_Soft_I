package model

import (
	"time"
)

type Curso struct {
	ID           int       `gorm:"primaryKey"`
	Fecha_Inicio time.Time `gorm:"type:date"`
	Fecha_Fin    time.Time `gorm:"type:date"`
	Materia_id   int       `gorm:"foreignKey"`
	Requisitos   string    `gorm:"type:varchar(255)"`
	Instructor   string    `gorm:"type:varchar(255)"`
}

type Cursos []Curso
