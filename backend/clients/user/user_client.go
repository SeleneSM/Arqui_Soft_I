package clients

import (
	"Arqui_Soft_I/model" //Llama a model
	"fmt"

	"github.com/jinzhu/gorm" //Importa el gorm
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByUsername(Username string) (model.User, error) {
	var user model.User
	result := Db.Where("username = ?", Username).First(&user)
	if result.Error != nil {
		return model.User{}, fmt.Errorf("error finding user by username: %w", result.Error)
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUsers() (model.Users, error) {
	var users model.Users
	result := Db.Find(&users)
	if result.Error != nil {
		return model.Users{}, fmt.Errorf("Error findig users: %w", result.Error)
	}
	log.Debug("User: ", users)
	return users, nil
}

func RegisterUser(user model.User) (model.User, error) {
	result := Db.Create(&user)

	if result.Error != nil {
		return model.User{}, fmt.Errorf("error registering user: %w", result.Error) //no hablo ingles
	}
	log.Debug("User Created: ", user.ID)
	return user, nil
}

func GetUserById(id int) (model.User, error) {
	var user model.User
	result := Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return model.User{}, fmt.Errorf("error geting user by id: %w", result.Error)
	}
	log.Debug("User: ", user)

	return user, nil
}
