package worker

import (
	"BangBot/api/botapi"
	"BangBot/config/botconfig"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type Task interface {
	Exec()
}

type Echo struct {
	Msg botapi.MessageNew
}

func (e *Echo) Exec() {
	fmt.Println(e.Msg)
}

type WriteMe struct{}

func (t *WriteMe) Exec() {
	params := url.Values{}
	params.Add("access_token", botconfig.ACCESSTOKEN)
	params.Add("message", "My king!")
	params.Add("random_id", strconv.Itoa(rand.Int()))
	params.Add("user_id", botconfig.MYVKID)
	params.Add("v", botconfig.VKVERSION)

	u, err := url.Parse("https://api.vk.com/method/messages.send")
	if err != nil {
		log.Fatal(err)
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

	fmt.Println(resp)
}
