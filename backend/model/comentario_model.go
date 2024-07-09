package model

import (
	"time"
)

type Comentario struct {
	ID               int       `gorm:"primaryKey"`
	ID_usuario       int       `gorm:"not null"`
	ID_curso         int       `gorm:"not null"`
	Usuario          User      `gorm:"foreignKey:UsuarioID;references:ID"`
	Fecha_comentario time.Time `gorm:"type:date"`
	Texto            string    `gorm:"type:text;not null"`
}

type Comentarios []Inscripcion
