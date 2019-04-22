package worker

import (
	"BangBot/api/botapi"
	"time"
)

var TasksBuffer = 10

var CallBackWorcker = NewWorker()

type Worker struct {
	Tasks chan Task
}

func NewWorker() *Worker {
	worker := &Worker{
		Tasks: make(chan Task, TasksBuffer),
	}

	go worker.Run()

	return worker
}

func (w *Worker) AddTask(task Task) {
	w.Tasks <- task
}

func (w *Worker) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.Tasks <- &Echo{
				Msg: botapi.MessageNew{
					Id: 10,
				},
			}

		case task := <-w.Tasks:
			task.Exec()
		}

	}
}
