package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type Users []User
