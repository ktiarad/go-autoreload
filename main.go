package main

import (
	"fmt"
	"go-autoreload/services"
	"net/http"
)

func main() {
	port := ":4444"
	http.HandleFunc("/", services.StartWeather)
	fmt.Println("Application is listening on port", port)
	http.ListenAndServe(port, nil)

}
