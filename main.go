package main

import (
	app "Login/app"
	"Login/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
