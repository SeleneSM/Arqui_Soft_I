package controller

import (
	"Arqui_Soft_I/backend/dto"
	service "Arqui_Soft_I/backend/service"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func InsertComentario(c *gin.Context) {
	var comentario dto.ComentarioDto

	if err := c.ShouldBindJSON(&comentario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debug("user_id: ", comentario.UsuarioID)
	comentario, apiErr := service.InsertComentario(comentario)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, comentario)
}

func GetComentariosPorCursoID(c *gin.Context) {
	cursoID := c.Param("curso_id")

	comentarios, apiErr := service.GetComentariosPorCursoID(cursoID)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, comentarios)
}
