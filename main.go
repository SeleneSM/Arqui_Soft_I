package main

import (
	app "Arqui_Soft_I/backend/app"
	"Arqui_Soft_I/backend/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
