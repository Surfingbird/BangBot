package worker

import (
	"BangBot/api/botapi"
	"fmt"
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
