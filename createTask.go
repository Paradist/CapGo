package capGo


func (c capGoStruct) CreateTask(task map[string]any) (*CapSolverResponse, error) {
	resp, err := c.Request(CREATE_TASK_PATH, &CapSolverRequest{
		ClientKey: c.ApiKey,
		Task:      &task,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}