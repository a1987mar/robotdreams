package main

import (
	"fmt"
	"net/http"
	"os"
	"robotlesson/config"
	"robotlesson/idternal/login"
)

func main() {
	config := config.LoadConfig()
	router := http.NewServeMux()
	login.NewLoginHandler(router)
	fmt.Println("Start server")
	server := http.Server{
		Addr:    config.Serv.Port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("STOP SERVER", err)
		os.Exit(2)
	}
	os.Exit(2)
}
