package app

import (
	userController "Login/controller/user"
)

// Endpoints
func mapUrls() {
	router.GET("users/auth", userController.UserAuth)
}
