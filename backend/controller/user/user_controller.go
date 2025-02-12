package controller

import (
	"Arqui_Soft_I/dto"
	jwtToken "Arqui_Soft_I/jwt"
	service "Arqui_Soft_I/service"
	"net/http"
	"strconv"

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

func GetUsers(c *gin.Context) {
	var usersDto dto.Users
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserRegister(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.RegisterUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"user_created": "true",
	})
}

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userDto)
}
