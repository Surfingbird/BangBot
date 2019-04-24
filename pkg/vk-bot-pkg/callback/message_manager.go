package callback

import (
	"BangBot/api/botapi"
	"BangBot/pkg/vk-bot-pkg/worker"
)

func msgLogic(msg botapi.MessageNew) {

	switch text := msg.Text; text {
	case "/help":
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})

	case "hellow":
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})

	default:
		worker.CallBackWorcker.AddTask(&worker.Hellow{
			Msg: msg,
		})
	}
}
