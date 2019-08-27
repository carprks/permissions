package service

import (
	"encoding/json"
)

func delete(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", err
	}

	resp, err := p.DeleteEntry()
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
