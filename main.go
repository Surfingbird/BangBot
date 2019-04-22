package main

import (
	"log"
	"net/http"

	"BangBot/api/botapi"
	"BangBot/config/botconfig"

	"github.com/gin-gonic/gin"
)

func callbackAnswer(c *gin.Context) {
	accept := &botapi.VKAccept{}
	err := c.BindJSON(accept)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	// if accept.Type != "confirmation" || accept.GroupId != 181466832 {
	// 	c.AbortWithStatus(http.StatusForbidden)

	// 	return
	// }

	log.Println(accept)
	c.String(http.StatusOK, "fa1fc205")
}

func CorsMiddlewareGin(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	c.Next()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CorsMiddlewareGin)

	router.POST("/callback", callbackAnswer)

	return router
}

func main() {
	log.Println(botconfig.BOTPORT)

	router := setupRouter()
	router.Run(":" + botconfig.BOTPORT)
}
