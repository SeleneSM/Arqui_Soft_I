package clients

import (
	"Arqui_Soft_I/backend/model" //Llama a model

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

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("User: ", users)
	return users
}

func RegisterUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("User Created: ", user.ID)
	return user
}

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}
