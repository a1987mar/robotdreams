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

	http.HandleFunc("GET /login", Log())

	fmt.Println("Start server")
	server := http.Server{
		Addr: config.Serv.Port,
	}
	server.ListenAndServe()

}

func Log() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var log login.Login
		json.NewDecoder(r.Body).Decode(&log)
		pkg.JsonResponse(w, log, 201)
	}

}
