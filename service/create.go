package service

import (
	"encoding/json"
	"fmt"
)

func create(body string) (string, error) {
	fmt.Println(fmt.Sprintf("create permissions start"))

	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", fmt.Errorf("create unmarshall err: %w", err)
	}

	if len(p.Permissions) == 0 {
		return "", fmt.Errorf("need at least 1 permission")
	}

	resp, err := p.CreateEntry()
	if err != nil {
		return "", fmt.Errorf("create entry err: %w", err)
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("create entry marshall err: %w", err)
	}

	fmt.Println(fmt.Sprintf("create permissions: %s", resp.Identifier))

	return string(res), nil
}
