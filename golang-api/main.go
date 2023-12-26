package main

import (
	"log"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok!"))
	})

	log.Print("server running on port :8000")
	_ = http.ListenAndServe(":8000", app)
}
