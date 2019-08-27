package service

import (
	"encoding/json"
)

func update(body string) (string, error) {
	req := Permissions{}
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", err
	}

	p := Permissions{
		Identifier: req.Identifier,
	}
	p, err = p.RetrieveEntry()
	if err != nil {
		return "", err
	}

	n, err := p.update(req)
	if err != nil {
		return "", err
	}

	resp, err := p.UpdateEntry(n)
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

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