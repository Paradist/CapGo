package capGo

import (
	"log"
	"time"
)

func (c capGoStruct) Solve(task map[string]any) (*CapSolverResponse, error) {
	log.Println("1")
	err := c.CheckTaskArguments(task)
	if err != nil {
		return nil, err
	}
	log.Println("1")
	resp, err := c.CreateTask(task)
	if err != nil {
		return nil, err
	}
	if (resp.Status == STATUS_READY) {
		return resp, nil
	}
	log.Println("1")
	for i := 0; i < TASK_TIMEOUT; i++ {
		log.Println("32")
		time.Sleep(time.Second)
		ok, _ := c.CheckTask(resp.TaskId)
		if (ok) {
			return resp, nil
		}
	}
	return nil, nil
}