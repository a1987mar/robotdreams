package main

import (
	"encoding/json"
	"net/http"
	"robotlesson/config"
)

func main() {
	config := config.LoadConfig()
	http.Handle("GET /login", Login())
	server := http.Server{
		Addr: config.Serv.Port,
	}
	server.ListenAndServe()
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world")
	}
}
