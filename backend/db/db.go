package db

import (
	cursoClient "Arqui_Soft_I/backend/clients/curso"
	incripcionClient "Arqui_Soft_I/backend/clients/inscripcion"
	materiaClient "Arqui_Soft_I/backend/clients/materia"
	userClient "Arqui_Soft_I/backend/clients/user"

	"Arqui_Soft_I/backend/model"
	"time"

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

	materias := []model.Materia{
		{
			Nombre:         "Python",
			Duracion:       3,
			Descripcion:    "Curso online de Python para principiantes. Aprende desde lo básico hasta conceptos avanzados, incluyendo estructuras de datos, funciones, y programación orientada a objetos, a través de ejercicios prácticos y proyectos reales.",
			Palabras_clave: "Programacion, Phyton, software, desarrollo",
		},
		{
			Nombre:         "C++",
			Duracion:       2,
			Descripcion:    "Curso online de C++ para todos los niveles. Domina desde la sintaxis básica hasta la programación avanzada, incluyendo manejo de memoria, POO, y algoritmos eficientes, a través de ejemplos prácticos y proyectos desafiantes",
			Palabras_clave: "C++, Sintaxis, Algoritmo, Memoria",
		},
		{
			Nombre:         "R",
			Duracion:       4,
			Descripcion:    "Curso online de R para análisis de datos. Aprende desde los fundamentos de R hasta técnicas avanzadas de visualización y modelado estadístico, aplicando tus conocimientos en proyectos prácticos y estudios de caso reales",
			Palabras_clave: "Analisis de datos, estadistica, modelado, visualizacion",
		},
	}

	for _, materia := range materias {
		var existingMateria model.Materia
		if err := db.Where("Nombre = ?", materia.Nombre).First(&existingMateria).Error; err == nil {
			log.Info("Materia already exists with Nombre:", materia.Nombre)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("Failed to query materia:", err.Error())
			continue
		}
		if err := db.Create(&materia).Error; err != nil {
			log.Error("Failed to insert materia:", err.Error())
		}
	}

	cursos := []model.Curso{
		{
			Fecha_Inicio: time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC),
			Fecha_Fin:    time.Date(2024, time.June, 30, 0, 0, 0, 0, time.UTC),
			Materia_id:   4,
			Requisitos:   "Disponibilidad de una computadora con conexion a internet 3 horas a la semana",
			Instructor:   "Matias Tossetti",
		},
		{
			Fecha_Inicio: time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC),
			Fecha_Fin:    time.Date(2024, time.October, 31, 0, 0, 0, 0, time.UTC),
			Materia_id:   5,
			Requisitos:   "Disponibilidad de una computadora con conexion a internet 5 horas a la semana",
			Instructor:   "Julio Gutierrez",
		},
		{
			Fecha_Inicio: time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC),
			Fecha_Fin:    time.Date(2024, time.September, 31, 0, 0, 0, 0, time.UTC),
			Materia_id:   4,
			Requisitos:   "Disponibilidad de una computadora con conexion a internet 3 horas a la semana",
			Instructor:   "Julio Gutierrez",
		},
	}

	for _, curso := range cursos {
		var existingCurso model.Curso
		if err := db.Where("Fecha_Inicio = ? AND Fecha_fin = ? AND Materia_id = ?", curso.Fecha_Inicio, curso.Fecha_Fin, curso.Materia_id).First(&existingCurso).Error; err == nil {
			log.Info("Curso already exists with id:", curso.Fecha_Inicio, curso.Fecha_Fin, curso.Materia_id)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("Failed to query curso:", err.Error())
			continue
		}
		if err := db.Create(&curso).Error; err != nil {
			log.Error("Failed to insert curso:", err.Error())
		}
	}

	/*inscripcion := model.Inscripcion{

		ID_curso:          1,
		ID_usuario:        2,
		Fecha_Inscripcion: time.Date(2024, time.September, 23, 0, 0, 0, 0, time.UTC),
	}

	if err := db.Create(&inscripcion).Error; err != nil {
		log.Error("Failed to insert curso:", err.Error())
	}*/

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
	cursoClient.Db = db
	materiaClient.Db = db
	incripcionClient.Db = db

}

func StartDbEngine() {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Materia{})
	db.AutoMigrate(&model.Curso{})
	db.AutoMigrate(&model.Inscripcion{})

	insertInitialData()

	log.Info("Finishing Migration Database Tables")
}
