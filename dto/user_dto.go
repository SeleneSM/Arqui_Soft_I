package dto

type UserDto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type Users []UserDto
