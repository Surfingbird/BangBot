package callback

import (
	"BangBot/api/botapi"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func CallbackAnswer(c *gin.Context) {
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(c.Request.Body)
	// s := buf.String()

	// fmt.Println(s)

	accept := &botapi.VKMsg{}
	err := c.BindJSON(accept)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	defer c.String(http.StatusOK, "ok")

	if accept.Type == "message_new" {
		msg := botapi.MessageNew{}
		err = mapstructure.Decode(accept.MSG, &msg)
		if err != nil {
			log.Println(err.Error())

			return
		}
		fmt.Println(msg.FromId)

		newmsg := &botapi.VKNewMsg{}
		err = c.BindJSON(newmsg)
		if err != nil {
			fmt.Println("не получилось")
		} else {
			fmt.Printf("Структура: %v\n id: %v", newmsg, newmsg.MSG.FromId)
		}

		msgLogic(msg)
	}
}
