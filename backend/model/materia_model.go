package model

type Materia struct {
	ID             int    `gorm:"primaryKey"`
	Nombre         string `gorm:"type:varchar(255);not null"`
	Duracion       int    `gorm:"type:int"`
	Descripcion    string `gorm:"type:varchar(255)"`
	Palabras_clave string `gorm:"type:varchar(255)"`
}

type Materias []Materia
