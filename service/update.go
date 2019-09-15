package service

import (
	"encoding/json"
	"fmt"
)

func update(body string) (string, error) {
	fmt.Println(fmt.Sprintf("update permissions start"))

	req := Permissions{}
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", fmt.Errorf("update unmarshall err: %w", err)
	}

	p := Permissions{
		Identifier: req.Identifier,
	}
	p, err = p.RetrieveEntry()
	if err != nil {
		return "", fmt.Errorf("update retrieve entry err: %w", err)
	}

	n, err := p.update(req)
	if err != nil {
		return "", fmt.Errorf("update update err: %w", err)
	}

	resp, err := p.UpdateEntry(n)
	if err != nil {
		return "", fmt.Errorf("update update entry err: %w", err)
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("update marsahll err: %w", err)
	}

	fmt.Println(fmt.Sprintf("update permissions: %s", resp.Identifier))
	return string(res), nil
}

func (p Permissions) update(n Permissions) (Permissions, error) {
	r := Permissions{
		Identifier: n.Identifier,
	}

	for _, oldPerm := range p.Permissions {
		for _, newPerm := range n.Permissions {
			if newPerm.Name == oldPerm.Name {
				if newPerm.Action == oldPerm.Action {
					np := Permission{
						Action:     newPerm.Action,
						Name:       newPerm.Name,
						Identifier: newPerm.Identifier,
					}

					r.Permissions = append(r.Permissions, np)
				} else {
					r.Permissions = append(r.Permissions, oldPerm)
				}
			} else {
				r.Permissions = append(r.Permissions, oldPerm)
			}
		}
	}

	return r, nil
}

func deity(body string) (string, error) {
	fmt.Println(fmt.Sprintf("deity permissions start"))

	req := Permissions{}

	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", fmt.Errorf("deity update unmarshall: %w", err)
	}

	p := Permissions{
		Identifier: req.Identifier,
	}

	p, err = p.RetrieveEntry()
	if err != nil {
		return "", fmt.Errorf("deity retrieve: %w", err)
	}

	req.Permissions = []Permission{
		{
			Name:       "*",
			Identifier: "*",
			Action:     "*",
		},
	}

	resp, err := p.UpdateEntry(req)
	if err != nil {
		return "", fmt.Errorf("deity update entry err: %w", err)
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("deity marsahll err: %w", err)
	}

	fmt.Println(fmt.Sprintf("deity permisions: %s", resp.Identifier))
	return string(res), nil
}
