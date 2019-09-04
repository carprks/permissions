package service

import (
	"encoding/json"
	"fmt"
)

func delete(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", fmt.Errorf("delete unmarsahll err: %w", err)
	}

	resp, err := p.DeleteEntry()
	if err != nil {
		return "", fmt.Errorf("delete entry err: %w", err)
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("delete marshall err: %w", err)
	}

	return string(res), nil
}
