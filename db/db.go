package db

import (
	userClient "Arqui_Soft_I/clients/user"
	"Arqui_Soft_I/model"

	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	db  *gorm.DB
	err error
)

func insertInitialData() {
	users := []model.User{
		{
			Username: "larapereyra",
			Password: "1111",
			Rol:      "Estudiante",
			Nombre:   "Lara",
			Apellido: "Pereyra",
		},
		{
			Username: "selenesaad",
			Password: "2222",
			Rol:      "Estudiante",
			Nombre:   "Selene",
			Apellido: "Saad",
		},
		{
			Username: "luanarolon",
			Password: "3333",
			Rol:      "Administrador",
			Nombre:   "Luana",
			Apellido: "Rolon",
		},
		{
			Username: "luciaangiolini",
			Password: "4444",
			Rol:      "Estudiante",
			Nombre:   "Lucia",
			Apellido: "Angiolini",
		},
	}
	for _, user := range users {
		var existingUser model.User
		if err := db.Where("Username = ?", user.Username).First(&existingUser).Error; err == nil {
			log.Info("User already exists with Username:", user.Username)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("Failed to query user:", err.Error())
			continue
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("Error al hashear la password:", err.Error())
			continue
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			log.Error("Failed to insert user:", err.Error())
		}
	}

	log.Info("Initial values inserted")

}

func init() {
	DBName := "cursify"
	DBUser := "root"
	DBPass := "Luchiucc2024."
	DBHost := "127.0.0.1"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
}

func StartDbEngine() {

	db.AutoMigrate(&model.User{})

	insertInitialData()

	log.Info("Finishing Migration Database Tables")
}
