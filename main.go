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
	r.HandleFunc("/health", healthPage)
	r.HandleFunc("/hello", helloPage)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func helloWorld() string {
	return "Hello World"
}

func health() map[string]bool {
	return map[string]bool{"ok": true}
}

func healthPage(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, health())
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	resp := helloWorld()
	respondWithJSON(w, http.StatusOK, resp)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{}
	response["data"] = payload
	json.NewEncoder(w).Encode(response)
}
