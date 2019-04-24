package callback

import (
	"BangBot/api/botapi"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func getFromId(str string) string {
	restr := regexp.MustCompile(`"from_id":[0-9]+`)
	pair := restr.Find([]byte(str))
	reid := regexp.MustCompile(`[0-9]+`)
	id := reid.Find(pair)

	return string(id)
}

func CallbackAnswer(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	s := buf.String()
	fmt.Println(s)

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

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		s := buf.String()

		fromId := getFromId(s)
		fmt.Printf("fromId: %v\n\n", fromId)
		msg.FromId = fromId
		fmt.Println(msg)

		msgLogic(msg)
	}
}
