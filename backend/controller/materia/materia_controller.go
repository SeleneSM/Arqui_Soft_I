package controller

import (
	"Arqui_Soft_I/dto"
	service "Arqui_Soft_I/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetMateriaById(c *gin.Context) {
	log.Debug("Materia id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var materiaDto dto.MateriaDto

	materiaDto, err := service.MateriaService.GetMateriaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, materiaDto)
}
func SearchMateria(c *gin.Context) {
	var materiasDto dto.Materias

	palabras_clave := c.Param("palabras_clave")
	log.Debug("Materia keyword to load: " + c.Param("palabras_clave"))

	materiasDto, err := service.MateriaService.SearchMateria(palabras_clave)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, materiasDto)

}

func MateriaInsert(c *gin.Context) {
	var materiaDto dto.MateriaDto

	err := c.BindJSON(&materiaDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	materiaDto, er := service.MateriaService.InsertMateria(materiaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, materiaDto)
}
