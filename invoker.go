package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvokeWebMethod(url string, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return ioutil.ReadAll(resp.Body)
	}
	return nil, errors.New(fmt.Sprint("status code is", resp.StatusCode))
}
