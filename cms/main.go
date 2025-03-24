package main

import (
	"net/http"
	"robotlesson/config"
	"robotlesson/pkg"
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
		pkg.JsonResponse(w, "HelloWorld", http.StatusOK)
	}
}
