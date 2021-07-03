package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":9001"
func main() {
	fmt.Printf("starting application on port %v", PORT)
	r := mux.NewRouter()
	r.HandleFunc("/health", landingPage)
	r.HandleFunc("/hello", helloPage)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func helloWorld() string {
	return "Hello World"
}

func health() map[string]bool {
	return map[string]bool {"ok": true}
}

func landingPage(w http.ResponseWriter, r * http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health())
}

func helloPage(w http.ResponseWriter, r * http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{}
	response["data"] = helloWorld()
	json.NewEncoder(w).Encode(response)
}