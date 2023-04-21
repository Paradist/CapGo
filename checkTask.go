package capGo

import "log"

func (c capGoStruct) CheckTask(taskId string) (bool, error) {
	resp, err := c.Request(GET_TASK_PATH, &CapSolverRequest{
		ClientKey: c.ApiKey,
		TaskId:    taskId,
	})
	log.Println(err)
	if err != nil {
		return false, err
	}
	log.Println(resp.Status)
	return resp.Status == STATUS_READY, nil
}