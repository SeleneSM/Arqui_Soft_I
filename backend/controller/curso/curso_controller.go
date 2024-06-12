package controller

import (
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

	c.JSON(http.StatusOK, cursos)
}
