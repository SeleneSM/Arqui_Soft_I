package controller

import (
	materiaClient "Arqui_Soft_I/backend/clients/materia"
	service "Arqui_Soft_I/backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCursos(c *gin.Context) {
	cursos, err := service.CursoService.GetCursos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message()})
		return
	}

	var result []gin.H // en esta lista se combina la informacion de cursos y su materia
	for _, curso := range cursos {
		// Buscar la materia por su ID en el dto de curso
		materia, err := materiaClient.GetMateriaById(curso.Materia_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get materia for curso"})
			return
		}

		// Combinar curso y materia en la respuesta
		result = append(result, gin.H{
			"id":           curso.ID,
			"fecha_inicio": curso.Fecha_Inicio,
			"fecha_fin":    curso.Fecha_Fin,
			"materia_id":   curso.Materia_id,
			"nombre":       materia.Nombre,
			"duracion":     materia.Duracion,
			"descripcion":  materia.Descripcion,
		})
	}

	c.JSON(http.StatusOK, result)
}
