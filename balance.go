package capGo


func (c capGoStruct) Balance() (float32, error) {
	resp, err := c.Request(BALANCE_PATH, &CapSolverRequest{
		ClientKey: c.ApiKey,
	})
	if err != nil {
		return 0, err
	}
	return resp.Balance, nil
}