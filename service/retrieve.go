package service

import (
	"encoding/json"
)

func retrieve(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", err
	}

	resp, err := p.RetrieveEntry()
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func allowed(body string) (string, error) {
	p := Permissions{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return "", err
	}

	resp, err := p.RetrieveEntry()
	if err != nil {
		return "", err
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
		return "", err
	}

	return string(res), nil
}
