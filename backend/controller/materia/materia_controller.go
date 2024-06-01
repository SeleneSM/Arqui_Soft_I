package controller

import (
	"Arqui_Soft_I/backend/dto"
	service "Arqui_Soft_I/backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
	func GetHotels(c *gin.Context) {
		var hotelsDto dto.HotelsDto
		hotelsDto, err := service.HotelService.GetHotels()

		if err != nil {
			c.JSON(err.Status(), err)
			return
		}

		c.JSON(http.StatusOK, hotelsDto)
	}
*/
func SearchMateria(c *gin.Context) {
	var materiasDto dto.Materias
	log.Debug("Materia keyword to load: " + c.Param("palabras_clave"))

	palabras_clave := c.Param("palabras_clave")
	materiasDto, err := service.MateriasService.SearchMateria(palabras_clave)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, materiasDto)

}
