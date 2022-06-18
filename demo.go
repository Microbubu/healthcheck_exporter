package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DemoApiCaller struct{}

type DemoApiResult struct {
	Data string `json:"d"`
}

func (caller *DemoApiCaller) Call(req *CallerRequest) *CallerResult {
	fmt.Println("DemoApiCaller called at", time.Now())
	body, err := InvokeWebMethod(req.Url, req.HttpMethod)
	if err != nil {
		return &CallerResult{}
	}
	if len(body) > 0 {
		res := &DemoApiResult{}
		if json.Unmarshal(body, res) != nil {
			return &CallerResult{}
		}
		fmt.Println(res.Data)
	}
	return &CallerResult{}
}
