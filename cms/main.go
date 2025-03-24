package main

import (
	"net/http"
	"robotlesson/config"
)

func main() {
	config := config.LoadConfig()

	server := http.Server{
		Addr: config.Serv.Port,
	}
	server.ListenAndServe()

}
