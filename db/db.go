package db

import (
	userClient "Arqui_Soft_I/clients/user"
	"Arqui_Soft_I/model"

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
	user := model.User{
		Username: "larapereyra",
		Password: "1111",
		Rol:      "Estudiante",
		Nombre:   "Lara",
		Apellido: "Pereyra",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error al hashear la password:", err.Error())
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		log.Error("Failed to insert user:", err.Error())
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
