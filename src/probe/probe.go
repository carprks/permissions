package probe

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP responds with a healthy model
func HTTP(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		log.Println("Probe Request", string(buf))
	}

	// get probe response
	resp, _ := Probe()

	// send status
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/health+json")
	w.Header().Set("Service", "Permissions")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	return
}

// Probe responds with a healthy model
func Probe() (Healthy, error) {
	return Healthy{
		Status: "pass",
	}, nil
}
