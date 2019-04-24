package main

import (
	"BangBot/config"
	"BangBot/config/botconfig"
	"BangBot/pkg/public/middleware"
	"BangBot/pkg/vk-bot-pkg/callback"
	"BangBot/pkg/vk-bot-pkg/token"
	"BangBot/pkg/vk-bot-pkg/worker"
	_ "BangBot/pkg/vk-bot-pkg/worker"
	"fmt"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddlewareGin)

	router.POST("/callback", callback.CallbackAnswer)
	router.POST("/token", token.UpdateTokenHandle)
	router.POST("/msgme", writeme)

	return router
}

func writeme(c *gin.Context) {
	worker.CallBackWorcker.AddTask(&worker.WriteMe{})
}

func main() {
	config.Logger.Info(fmt.Sprintf("BOTPORT: %v", botconfig.BOTPORT))
	config.Logger.Info(fmt.Sprintf("VKAPPID: %v", botconfig.VKAPPID))
	config.Logger.Info(fmt.Sprintf("SECUREKEY: %v", botconfig.SECUREKEY))
	config.Logger.Info(fmt.Sprintf("VKVERSION: %v", botconfig.VKVERSION))

	defer config.Logger.Sync()

	router := setupRouter()
	router.Run(":" + botconfig.BOTPORT)
}
