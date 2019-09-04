package service

import (
	"encoding/json"
	"fmt"
)

func retrieve(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", fmt.Errorf("retieve unmarshall err: %w", err)
	}

	resp, err := p.RetrieveEntry()
	if err != nil {
		return "", fmt.Errorf("retrieve entry err: %w", err)
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("retrieve marshall err: %w", err)
	}

	return string(res), nil
}

func allowed(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", fmt.Errorf("allowed unmarshall err: %w", err)
	}

	resp, err := p.RetrieveEntry()
	if err != nil {
		return "", fmt.Errorf("allowed retrieve entry err: %w", err)
	}

	allowed := false
	for _, perm := range resp.Permissions {
		for _, cperm := range p.Permissions {
			if perm.Name == cperm.Name {
				if perm.Action == cperm.Action {
					if perm.Identifier == p.Identifier || perm.Identifier == "*" {
						allowed = true
					}
				}

				if perm.Action == "*" {
					allowed = true
				}
			}

			if perm.Name == "*" {
				allowed = true
			}
		}
	}

	if allowed {
		resp = Permissions{
			Identifier: p.Identifier,
			Status:     "allowed",
		}
	} else {
		resp = Permissions{
			Identifier: p.Identifier,
			Status:     "denied",
		}
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("allowed marshall err: %w", err)
	}

	return string(res), nil
}
