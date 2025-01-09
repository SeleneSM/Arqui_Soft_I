package dto

type ArchivoDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Curso_id int    `json:"curso_id"`
}

type ArchivosDto []ArchivoDto
