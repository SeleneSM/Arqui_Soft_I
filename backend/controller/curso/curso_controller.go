package controller

import (
	"Arqui_Soft_I/backend/dto"
	service "Arqui_Soft_I/backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCursos(c *gin.Context) {
	var cursosDto dto.Cursos
	cursosDto, err := service.CursoService.GetCursos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, cursosDto)
}
