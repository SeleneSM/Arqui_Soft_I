package clients

import (
	"Arqui_Soft_I/model" //Llama a model

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByUsername(Username string) model.User {
	var user model.User
	Db.Where("username = ?", Username).First(&user)
	log.Debug("User: ", user)

	return user
}
