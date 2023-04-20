package capGo

import "time"

func (c capGoStruct) Solve(task map[string]any) (*CapSolverResponse, error) {
	err := c.CheckTaskArguments(task)
	if err != nil {
		return nil, err
	} 
	resp, err := c.CreateTask(task)
	if err != nil {
		return nil, err
	}
	for i := 0; i < TASK_TIMEOUT; i++ {
		time.Sleep(time.Second)
		ok, err := c.CheckTask(resp.TaskId)
		if (err != nil) {
			return nil, err
		}
		if (ok) {
			return resp, nil
		}
	}
	return nil, nil
}