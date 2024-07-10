package app

import (
	comentarioController "Arqui_Soft_I/backend/controller/comentario"
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
	router.GET("/comentarios/curso/:curso_id", comentarioController.GetComentariosPorCursoID)

	//Post
	router.POST("/inscribir", inscripcionController.InscribirUsuario)
	router.POST("/users/auth", userController.UserAuth)
	router.POST("/registrar_usuario", userController.UserRegister)
	router.POST("/crear_materia", materiaController.MateriaInsert)
	router.POST("/crear_curso", cursoController.CursoInsert)
	router.POST("/comentarios", comentarioController.InsertComentario)

	//router.POST("cursos", cursoController.)
}
