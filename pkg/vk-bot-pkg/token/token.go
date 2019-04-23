package token

import (
	"github.com/gin-gonic/gin"
)

type accessToken struct {
	Token string `json:"token"`
}

// сомневаюсь, что так стоит делать
func UpdateTokenHandle(c *gin.Context) {
	// token := &accessToken{}
	// err := c.BindJSON(token)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusBadRequest)

	// 	return
	// }

	// // небезопасно, просто поиграться
	// botconfig.ACCESSTOKEN = token.Token
}
