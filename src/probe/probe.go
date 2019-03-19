package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Probe responds with a healthy model
func Probe(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		log.Println("Probe Request", string(buf))
	}

	// send status
	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(Healthy{
		Status: "pass",
	})
	w.Write(j)
	w.WriteHeader(http.StatusOK)

	return
}

func Tester(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		log.Println("Probe Request", string(buf))
	}

	// send status
	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(Healthy{
		Status: "pass",
	})
	w.Write(j)
	w.WriteHeader(http.StatusOK)

	fmt.Printf("Tester \n")

	return
}