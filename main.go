package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/otel"
)

func main() {
	var PORT = getPort()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("hello-oh-my"),
		launcher.WithAccessToken(os.Getenv("LIGHTSTEP")),
	)
	defer ls.Shutdown()

	fmt.Printf("starting application on port %v", PORT)
	r := mux.NewRouter()
	r.HandleFunc("/health", healthPage)
	r.HandleFunc("/", healthPage)
	r.HandleFunc("/hello", helloPage)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func helloWorld() string {
	tracer := otel.Tracer("example")
	_, span := tracer.Start(context.Background(), "main")
	defer span.End()

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

// Get the Port from the environment so we can run on Heroku
func getPort() string {
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "9001"
		fmt.Println("no port set in environment variable, using default port " + PORT)
	}
	return ":" + PORT
}
