package model

type Archivo struct {
	Id    int    `gorm:"primarykey"`
	Name  string `gorm:"type:varchar(200);not null"`
	Path  string `gorm:"type:varchar(200);not null"`
	Curso Curso  `gorm:"foreigkey:Curso_id"`

	Curso_Id int
}

type Archivos []Archivo
