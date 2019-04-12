package healthcheck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTP the request as done by routing
func HTTP(w http.ResponseWriter, r *http.Request) {
	hc := HealthCheck{
		Name: os.Getenv("SERVICE_NAME"),
		URL: r.Host,
		Dependencies: os.Getenv("SERVICE_DEPENDENCIES"),
	}

	health, err := hc.Check()
	if err != nil {
		w.Header().Set("Content-Type", "application/health+json")
		j, _ := json.Marshal(Health{
			Status: HealthFail,
		})
		w.Write(j)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(health)
	w.Write(j)
	w.WriteHeader(http.StatusOK)
}

// Check do the health check itself
func (h HealthCheck) Check() (Health, error) {
	health := Health{
		Name: h.Name,
		URL: h.URL,
		Status: HealthFail,
		Dependencies: nil,
	}

	if h.Dependencies != "" {
		deps, err := h.getDependencies()
		if err != nil {
			return health, err
		}

		checkedDeps := []Health{}
		for _, dep := range deps.Dependencies {
			d, err := dep.check()
			if err != nil {
				return health, err
			}
			checkedDeps = append(checkedDeps, d)
		}

		health.Dependencies = checkedDeps
	}
	health.Status = HealthPass

	return health, nil
}

// getDependencies get the list of dependencies
func (h HealthCheck) getDependencies() (Dependencies, error) {
	deps := Dependencies{}
	err := json.Unmarshal([]byte(h.Dependencies), &deps)
	if err != nil {
		return deps, err
	}

	return deps, nil
}

// check the dependency status
func (d Dependency) check() (Health, error) {
	// Ping check
	if d.Ping {
		return d.ping()
	}

	// Standard check
	return d.curl()
}

// ping checks
func (d Dependency) ping() (Health, error) {
	h := Health{
		Name: d.Name,
		URL: d.URL,
		Status: HealthFail,
	}

	p, err := http.Get(h.URL)
	if err != nil {
		return h, err
	}
	if p.StatusCode == http.StatusOK {
		h.Status = HealthPass
	}
	return h, nil
}

// curl checks
func (d Dependency) curl() (Health, error) {
	h := Health{}

	url := fmt.Sprintf("https://%s/healthcheck", d.URL)
	p, err := http.Get(url)
	if err != nil {
		h = Health{
			URL: url,
			Status: HealthFail,
		}
		return h, err
	}
	b, err := ioutil.ReadAll(p.Body)
	if err != nil {
		h = Health{
			URL: url,
			Status: HealthFail,
		}
		return h, err
	}
	jerr := json.Unmarshal(b, &h)
	if err != nil {
		return h, jerr
	}
	return h, nil
}