package capGo

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c capGoStruct) Request(path string, requestBody *CapSolverRequest) (*CapSolverResponse, error) {
	client := resty.New()

	body, err := json.Marshal(requestBody)
	if (err != nil) {
		return nil, err
	}

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).Get(fmt.Sprint(API_URL, path))
	if (err != nil) {
		return nil, err
	}

	responce := &CapSolverResponse{}
	responce.Solution = &Solution{}

	err = json.Unmarshal(resp.Body(), responce)
	if (err != nil) {
		return nil, err
	}
	return responce, nil
}