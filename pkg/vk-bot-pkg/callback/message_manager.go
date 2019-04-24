package callback

import (
	"BangBot/api/botapi"
	"BangBot/pkg/vk-bot-pkg/worker"
	"fmt"
)

// тут есть дебаг
func msgLogic(msg botapi.MessageNew) {
	fmt.Println(msg)
	fmt.Println(msg.FromId)

	switch text := msg.Text; text {
	case "/help":
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})
		fmt.Println("/help")

	case "hellow":
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})
		fmt.Println("hellow")

	default:
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})
		fmt.Println("default")

	}
}
