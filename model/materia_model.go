package model

type Materia struct {
	ID       int    `gorm:"primaryKey"`
	Nombre   string `gorm:"type:varchar(255);not null"`
	Duracion int    `gorm:"type:int;not null"`
}

type Materias []Materia
