package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		response := Response{"Hello world!"}
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}

