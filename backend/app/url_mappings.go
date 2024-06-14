package app

import (
	cursoController "Arqui_Soft_I/backend/controller/curso"
	inscripcionController "Arqui_Soft_I/backend/controller/inscripcion"
	materiaController "Arqui_Soft_I/backend/controller/materia"
	userController "Arqui_Soft_I/backend/controller/user"
)

// Endpoints
func mapUrls() {
	//Get

	router.GET("/materia/search/:palabras_clave", materiaController.SearchMateria)
	router.GET("/inscripciones_por_usuario/:id_usuario", inscripcionController.GetInscripcionesByUser)
	router.GET("/cursos", cursoController.GetCursos)
	router.GET("/materia/:id", materiaController.GetMateriaById)
	router.GET("/users", userController.GetUsers)

	//Post
	router.POST("/inscribir", inscripcionController.InscribirUsuario)
	router.POST("/users/auth", userController.UserAuth)
	//router.POST("cursos", cursoController.)
}
