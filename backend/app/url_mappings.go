package app

import (
	inscripcionController "Arqui_Soft_I/backend/controller/inscripcion"
	materiaController "Arqui_Soft_I/backend/controller/materia"
	userController "Arqui_Soft_I/backend/controller/user"
)

// Endpoints
func mapUrls() {
	router.GET("users/auth", userController.UserAuth)
	router.GET("materia/search/:palabras_clave", materiaController.SearchMateria)
	router.POST("inscribir", inscripcionController.InscribirUsuario)
	router.GET("inscripciones_por_usuario/:id_usuario", inscripcionController.GetInscripcionesByUser)
}
