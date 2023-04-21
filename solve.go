package capGo

import (
	"time"
)

func (c capGoStruct) Solve(task map[string]any) (*CapSolverResponse, error) {
	err := c.CheckTaskArguments(task)
	if err != nil {
		return nil, err
	}
	resp, err := c.CreateTask(task)
	if err != nil {
		return nil, err
	}
	if (resp.Status == STATUS_READY) {
		return resp, nil
	}
	for i := 0; i < TASK_TIMEOUT; i++ {
		time.Sleep(time.Second)
		ok, _ := c.CheckTask(resp.TaskId)
		if (ok) {
			return resp, nil
		}
	}
	return nil, nil
}