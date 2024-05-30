package app

import (
	userController "Arqui_Soft_I/controller/user"
)

// Endpoints
func mapUrls() {
	router.GET("users/auth", userController.UserAuth)
}
