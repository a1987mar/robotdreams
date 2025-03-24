package main

import (
	"net/http"
	"robotlesson/config"
)

func main() {
	config := config.LoadConfig()

	http.HandleFunc("POST /login", Login())

	server := http.Server{
		Addr: config.Serv.Port,
	}
	server.ListenAndServe()

}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}
