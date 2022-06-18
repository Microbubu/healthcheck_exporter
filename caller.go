package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Caller interface {
	Call(req *CallerRequest) *CallerResult
}

type CallerResult struct {
	Service     string
	IsUp        bool
	CallSuccess bool
	StatusCode  int
}

type CallerRequest struct {
	Url        string
	HttpMethod string
}

type EAPCaller struct{}
type MesCaller struct{}

func (caller *EAPCaller) Call(req *CallerRequest) *CallerResult {
	fmt.Println("EAPCaller called at", time.Now())
	return &CallerResult{
		Service:     "EAP",
		IsUp:        true,
		CallSuccess: true,
		StatusCode:  200,
	}
}

func (caller *MesCaller) Call(req *CallerRequest) *CallerResult {
	fmt.Println("MesCaller called at", time.Now())
	return &CallerResult{
		Service:     "EAP",
		IsUp:        true,
		CallSuccess: true,
		StatusCode:  200,
	}
}

func NewCaller(callerType string) (Caller, error) {
	switch strings.ToUpper(callerType) {
	case "EAP":
		return &EAPCaller{}, nil
	case "MES":
		return &MesCaller{}, nil
	case "DEMO":
		return &DemoApiCaller{}, nil
	default:
		return nil, errors.New("unknown caller type")
	}
}
