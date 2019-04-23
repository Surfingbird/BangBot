package callback

import (
	"BangBot/api/botapi"
	"log"
	"net/http"

	"BangBot/pkg/vk-bot-pkg/worker"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func CallbackAnswer(c *gin.Context) {
	accept := &botapi.VKMsg{}
	err := c.BindJSON(accept)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if accept.Type == "message_new" {
		msg := botapi.MessageNew{}
		err = mapstructure.Decode(accept.MSG, &msg)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusOK, "ok")

			return
		}

		// ToDo пока тут будет echo
		worker.CallBackWorcker.AddTask(&worker.Echo{
			Msg: msg,
		})
	}

	c.String(http.StatusOK, "ok")
}
