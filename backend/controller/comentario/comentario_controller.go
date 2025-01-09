package controller

import (
	"Arqui_Soft_I/dto"
	service "Arqui_Soft_I/service"
	"net/http"

	log "github.com/sirupsen/logrus"

	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertComentario(c *gin.Context) {
	var comentario dto.ComentarioDto

	if err := c.ShouldBindJSON(&comentario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debug("user_id: ", comentario.UsuarioID)
	log.Debug("curso_id: ", comentario.CursoID)
	comentario, apiErr := service.ComentarioService.InsertComentario(comentario)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, comentario)
}

func GetComentariosPorCursoID(c *gin.Context) {
	cursoIDStr := c.Param("curso_id")
	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid curso_id"})
		return
	}
	comentarios, apiErr := service.ComentarioService.GetComentariosPorCursoID(cursoID)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, comentarios)
}
