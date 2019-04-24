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

		fmt.Println("------Тут идет приведение типов------")
		msg1 := botapi.MessageNewInt{}
		err = mapstructure.Decode(accept.MSG, &msg1)
		if err != nil {
			log.Println(err.Error())

			return
		}

		fmt.Printf("int: %v", msg1.FromId)

		msg2 := botapi.MessageNewStr{}
		err = mapstructure.Decode(accept.MSG, &msg2)
		if err != nil {
			log.Println(err.Error())

			return
		}

		fmt.Printf("string: %v", msg2.FromId)

		msg3 := botapi.MessageNewFloat{}
		err = mapstructure.Decode(accept.MSG, &msg3)
		if err != nil {
			log.Println(err.Error())

			return
		}

		fmt.Printf("float: %v", msg3.FromId)
		fmt.Println("------Тут идет приведение типов------")

		// fmt.Printf("Пустой интерфейс: %v", msg.FromId)
		// idint, ok := (msg.FromId).(int)
		// if ok {
		// 	fmt.Printf("int: %v", idint)
		// }

		// idstr, ok := (msg.FromId).(string)
		// if ok {
		// 	fmt.Printf("string: %v", idstr)
		// }

		// idfloat, ok := (msg.FromId).(float64)
		// if ok {
		// 	fmt.Printf("float64: %v", idfloat)
		// }
		// fmt.Println("------Тут идет приведение типов------")

		msgLogic(msg)
	}
}
