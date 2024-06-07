package app

import (
	materiaController "Arqui_Soft_I/backend/controller/materia"
	userController "Arqui_Soft_I/backend/controller/user"
)

// Endpoints
func mapUrls() {
	router.GET("users/auth", userController.UserAuth)
	router.GET("materia/search/:palabras_clave", materiaController.SearchMateria)
}
