package dto

import (
	"time"
)

type ComentarioDto struct {
	ID               int       `json:"id"`
	CursoID          int       `json:"curso_id"`
	UsuarioID        int       `json:"usuario_id"`
	Texto            string    `json:"texto"`
	Fecha_comentario time.Time `json:"fecha"`
	Usuario          UserDto   `json:"usuario"`
}

type Comentarios []ComentarioDto
