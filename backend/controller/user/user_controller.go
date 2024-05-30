package controller

import (
	"Arqui_Soft_I/backend/dto"
	jwtToken "Arqui_Soft_I/backend/jwt"
	service "Arqui_Soft_I/backend/service"
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
	var rol string
	auth, rol, id = service.UserService.UserAuth(userDto)
	if auth {
		userDto.ID = id
		userDto.Rol = rol

		token, err := jwtToken.GenerateUserToken(userDto)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"auth":    true,
			"user_id": id,
			"rol":     rol,
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"auth":    false,
			"message": "Invalid email or password",
		})
	}
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
