package controller

import (
	"Arqui_Soft_I/dto"
	service "Arqui_Soft_I/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserAuth(c *gin.Context) {
	var userDto dto.UserDto

	err := c.BindJSON(&userDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var auth bool
	var id int
	auth, id = service.UserService.UserAuth(userDto)
	if auth == true {
		userDto.ID = id
	}

	c.JSON(http.StatusAccepted, gin.H{
		"auth":    auth,
		"user_id": id,
	})
}

/*func UserAuth(c *gin.Context) {
	var userDto dto.UserDto

	err := c.BindJSON(&userDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var autenticado bool
	var tipo int
	var id int
	autenticado, tipo, id = service.UserService.UserAuth(userDto)
	if autenticado == true {
		userDto.Tipo = tipo
		userDto.Id = id
		token, err := jwtG.GenerateUserToken(userDto)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"autenticacion": "true",
			"tipo":          tipo,
			"user_id":       id,
			"token":         token,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"autenticacion": "false",
			"tipo":          tipo,
			"user_id":       id,
		})
	}

}*/
