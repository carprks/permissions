package service

import (
	"encoding/json"
	"fmt"
)

func create(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", err
	}

	if len(p.Permissions) == 0 {
		err = fmt.Errorf("need at least 1 permission")
		return "", err
	}

	resp, err := p.CreateEntry()
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
