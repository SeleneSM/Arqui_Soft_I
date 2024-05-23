package dto

type CursoDto struct {
	ID           int `json:"id"`
	Fecha_Inicio int `json:"fecha_inicio"`
	Fecha_Fin    int `json:"fecha_fin"`
}

type Cursos []CursoDto
