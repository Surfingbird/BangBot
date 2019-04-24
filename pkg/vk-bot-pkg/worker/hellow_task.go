package worker

import (
	"BangBot/api/botapi"
	"BangBot/config/botconfig"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type Hellow struct {
	Msg botapi.MessageNew
}

func (t *Hellow) Exec() {
	// u, err := url.Parse("https://api.vk.com/method/messages.send")
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// q := u.Query()
	// q.Set("access_token", botconfig.ACCESSTOKEN)
	// q.Set("message", "Hellow!")
	// q.Set("random_id", strconv.Itoa(rand.Int()))
	// q.Set("user_id", strconv.Itoa(t.Msg.FromId))
	// q.Set("v", botconfig.VKVERSION)

	// u.RawQuery = q.Encode()
	// resp, err := http.Post(u.String(), "text/hmtl", nil)
	// if err != nil {
	// 	log.Println(err.Error())

	// 	return
	// }

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(resp.Body)
	// s := buf.String()

	// fmt.Println(s)

	u, err := url.Parse("https://api.vk.com/method/messages.send")
	if err != nil {
		log.Println(err.Error())
	}

	q := u.Query()
	q.Set("access_token", botconfig.ACCESSTOKEN)
	q.Set("message", "My king!")
	q.Set("random_id", strconv.Itoa(rand.Int()))
	q.Set("user_id", botconfig.MYVKID)
	q.Set("v", botconfig.VKVERSION)

	u.RawQuery = q.Encode()
	resp, err := http.Post(u.String(), "text/hmtl", nil)
	if err != nil {
		log.Println(err.Error())

		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()

	fmt.Println(s)
}
