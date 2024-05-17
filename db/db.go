package db

import (
	userClient "Login/clients/user"
	"Login/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func insertInitialData() {
	user := model.User{
		Username: "Admin",
		Password: "Admin",
	}
	if err := db.Create(&user).Error; err != nil {
		log.Error("Failed to insert user: ", err.Error())
	}

	log.Info("Initial values inserted")
}

func init() {
	DBName := "Login"
	DBUser := "root"
	DBPass := "root"
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
