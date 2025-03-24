package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"robotlesson/config"
	"robotlesson/internal/login"
	pkg "robotlesson/pkg/respos"
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
		var log login.Login
		body := json.NewDecoder(r.Body).Decode(&log)
		if body != nil {
			http.Error(w, "bad reqyest", http.StatusBadRequest)
		}
		pkg.JsonResponse(w, body, http.StatusOK)
	}

}
