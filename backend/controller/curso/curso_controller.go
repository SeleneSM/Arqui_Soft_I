package controller

import (
	service "Arqui_Soft_I/backend/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"Arqui_Soft_I/backend/dto"

	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetCursos(c *gin.Context) {
	cursos, err := service.CursoService.GetCursos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, cursos)
}

func GetCursoById(c *gin.Context) {
	log.Debug("Curso id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var cursoDto dto.CursoDto

	cursoDto, err := service.CursoService.GetCursoById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, cursoDto)
}

func CursoInsert(c *gin.Context) {
	var cursoDto dto.CursoDto

	err := c.BindJSON(&cursoDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cursoDto, er := service.CursoService.InsertCurso(cursoDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, cursoDto)
}
