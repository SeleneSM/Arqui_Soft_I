package controller

import (
	"Arqui_Soft_I/backend/dto"
	service "Arqui_Soft_I/backend/service"
	"net/http"

	log "github.com/sirupsen/logrus"

	"strconv"

	"github.com/gin-gonic/gin"
)

func InscribirUsuario(c *gin.Context) {
	var inscripcion dto.InscripcionDto
	//BindJSON hace en marshall
	if err := c.BindJSON(&inscripcion); err != nil {
		//Posible error en el marshall
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Debug("user_id", inscripcion.ID_usuario)

	inscripcion, er := service.InscripcionService.InscribirUsuario(inscripcion)

	if er != nil {
		//Posible error en el service
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusAccepted, inscripcion)
}

func GetInscripcionesByUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id_usuario"))
	log.Debug("USER_ID: ", userId)
	var inscripcionesDto dto.Inscripciones
	inscripcionesDto, err := service.InscripcionService.GetInscripcionesByUser(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, inscripcionesDto)
}
