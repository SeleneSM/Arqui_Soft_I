package app

import (
	archivoController "Arqui_Soft_I/controller/archivo"
	comentarioController "Arqui_Soft_I/controller/comentario"
	cursoController "Arqui_Soft_I/controller/curso"
	inscripcionController "Arqui_Soft_I/controller/inscripcion"
	materiaController "Arqui_Soft_I/controller/materia"
	userController "Arqui_Soft_I/controller/user"
)

// Endpoints
func mapUrls() {
	//Get

	router.GET("/materia/search/:palabras_clave", materiaController.SearchMateria)
	router.GET("/inscripciones_por_usuario/:id_usuario", inscripcionController.GetInscripcionesByUser)
	router.GET("/cursos", cursoController.GetCursos)
	router.GET("/materia/:id", materiaController.GetMateriaById)
	router.GET("/materias", materiaController.GetMaterias)
	//router.GET("/users", userController.GetUsers)
	router.GET("/comentarios/curso/:curso_id", comentarioController.GetComentariosPorCursoID)
	router.GET("/cursos/:id", cursoController.GetCursoById)
	router.GET("/users/:id", userController.GetUserById)

	//Post
	router.POST("/inscribir", inscripcionController.InscribirUsuario)
	router.POST("/users/auth", userController.UserAuth)
	router.POST("/registrar_usuario", userController.UserRegister)
	router.POST("/crear_materia", materiaController.MateriaInsert)
	router.POST("/crear_curso", cursoController.CursoInsert)
	router.POST("/comentarios", comentarioController.InsertComentario)
	// Aquí defines el endpoint y lo vinculas con tu controlador
	router.POST("/upload-file", archivoController.UploadFileHandler)
	//router.POST("cursos", cursoController.)
}
