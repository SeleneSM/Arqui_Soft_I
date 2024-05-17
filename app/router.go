package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}

	router.Use(cors.New(corsConfig))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting Server")
	router.Run(":8090")
}
