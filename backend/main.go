package main

import (
	app "Arqui_Soft_I/app"
	"Arqui_Soft_I/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
