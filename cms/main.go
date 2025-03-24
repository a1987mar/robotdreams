package main

import (
	"fmt"
	"net/http"
	"robotlesson/config"
	"robotlesson/pkg"
)

func main() {
	config := config.LoadConfig()
	http.Handle("GET /login", Login())
	fmt.Println("Start server")
	server := http.Server{
		Addr: config.Serv.Port,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Stop server")
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pkg.JsonResponse(w, "Hello World", http.StatusOK)
	}
}
