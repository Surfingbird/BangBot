package main

import (
	"log"

	"BangBot/config/botconfig"
	"BangBot/pkg/public/middleware"
	"BangBot/pkg/vk-bot-pkg/callback"
	_ "BangBot/pkg/vk-bot-pkg/worker"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddlewareGin)

	router.POST("/callback", callback.CallbackAnswer)

	return router
}

func main() {
	log.Println(botconfig.BOTPORT)

	router := setupRouter()
	router.Run(":" + botconfig.BOTPORT)
}
