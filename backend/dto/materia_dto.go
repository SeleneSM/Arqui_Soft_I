package dto

type MateriaDto struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Duracion int    `json:"duracion"`
}

type Materias []MateriaDto
