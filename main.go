package main

import (
	"Go_lang_rest_API/api"
	"log"
	"net/http"
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
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
