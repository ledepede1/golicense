package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	if !GetUseFivem() {
		fmt.Println("UsingFiveM: false")
		router.HandleFunc("/api/{license}/{ipadresss}", API).Methods("GET")
	} else {
		fmt.Println("UsingFiveM: true")
		router.HandleFunc("/api/{resourcename}/{license}/{ipadresss}", API).Methods("GET")
	}

	err := http.ListenAndServe(":"+GetPort(), router)
	if err != nil {
		fmt.Println(err)
	}
}
