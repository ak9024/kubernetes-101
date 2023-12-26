package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	Response struct {
		Env interface{} `json:"env"`
	}
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env := Response{
			Env: os.Getenv("MESSAGE"),
		}
		byteEnv, _ := json.Marshal(env)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(byteEnv)
	})

	app.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok!"))
	})

	log.Print("server running on port :8000")
	_ = http.ListenAndServe(":8000", app)
}
