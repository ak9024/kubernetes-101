package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	Response struct {
		Env map[string]interface{} `json:"env"`
	}
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env := map[string]interface{}{
			"MESSAGE":       os.Getenv("MESSAGE"),
			"SECRET_CONFIG": os.Getenv("SECRET_CONFIG"),
		}
		envResponse := Response{
			Env: env,
		}
		byteEnv, _ := json.Marshal(envResponse)

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
