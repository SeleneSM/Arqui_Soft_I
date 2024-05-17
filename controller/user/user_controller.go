package controller

import (
	"Login/dto"
	service "Login/service"
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
