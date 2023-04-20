package capGo

func (c capGoStruct) CheckTask(taskId string) (bool, error) {
	resp, err := c.Request(TASK_PATH, &CapSolverRequest{
		ClientKey: 	c.ApiKey,
		TaskId:     taskId,
	})
	if err != nil {
		return false, err
	}
	return resp.Status == STATUS_READY, nil
}