package dto

type MateriaDto struct {
	ID             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Duracion       int    `json:"duracion"`
	Descripcion    string `json:"descripcion"`
	Palabras_clave string `json:"palabras_clave"`
}

type Materias []MateriaDto
