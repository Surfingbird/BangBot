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

	if accept.Type != "confirmation" || accept.GroupId != 181466832 {
		c.AbortWithStatus(http.StatusForbidden)
		
		return
	}

	log.Println(accept)
	c.String(http.StatusOK, "fa1fc205")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/callback", callbackAnswer)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":" + botconfig.BOTPORT)
}
