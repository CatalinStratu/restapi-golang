package main

import (
	"fmt"
	"net/http"
	"service/api"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

func main() {

	mx := http.NewServeMux()
	mx.HandleFunc("/just/an/example/how_much", api.HowMuch)
	mx.HandleFunc("/just/an/example/list_how_many", api.ListHowMuch)
	mx.HandleFunc("/", api.ErrorHandler)
	err := http.ListenAndServe(Host+":"+Port, mx)
	if err != nil {
		fmt.Printf("Error Starting the HTTP Server: %v", err)
		return
	}
}
