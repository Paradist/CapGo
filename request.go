package capGo

import (
	"fmt"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func (c capGoStruct) Request(path string, requestBody *CapSolverRequest) (*CapSolverResponse, error) {
	client := resty.New()

	body, err := json.Marshal(requestBody)
	if (err != nil) {
		return nil, err
	}

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).Post(fmt.Sprint(API_URL, path))
	if (err != nil) {
		return nil, err
	}

	response := &CapSolverResponse{}
	response.Solution = &Solution{}

	err = json.Unmarshal(resp.Body(), response)
	if (err != nil) {
		return nil, err
	}
	if response.ErrorId == 1 {
		return nil, fmt.Errorf("%s: %s", response.ErrorCode, response.ErrorDescription)
	}
	
	return response, nil
}