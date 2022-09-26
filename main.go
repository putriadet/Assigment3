package main

import (
	"assigment3/service"
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	go service.AutoReload()
	http.HandleFunc("/", service.ReloadWeb)
	fmt.Println("Auto Reload App is listening on Port", port)
	http.ListenAndServe(port, nil)
}
