package capGo

func (c capGoStruct) Balance() (float32, error){
	resp, err := c.Request(BALANCE_URI, &CapSolverRequest{
		ClientKey: c.ApiKey,
	})
   
	return resp.Balance, err
}