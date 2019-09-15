package service

import (
	"encoding/json"
	"fmt"
)

func delete(body string) (string, error) {
	fmt.Println(fmt.Sprintf("delete permissions start"))

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

	fmt.Println(fmt.Sprintf("delete permissions: %s", resp.Identifier))

	return string(res), nil
}
