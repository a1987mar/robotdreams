package main

import (
	"fmt"
	"robotlesson/config"
)

func main() {
	config := config.LoadConfig()
	fmt.Printf("HELLO WORLD \nPort = %s", config.Serv.Port)
}
