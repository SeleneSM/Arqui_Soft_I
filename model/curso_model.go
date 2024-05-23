package model

type Curso struct {
	ID           int `gorm:"primaryKey"`
	Fecha_Inicio int `gorm:"type:int;not null"`
	Fecha_Fin    int `gorm:"type:int;not null"`
}

type Cursos []Curso
