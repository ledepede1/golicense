package main

import (
	"fmt"
	"net/http"

	"github.com/TwiN/go-color"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	if !GetUseFivem() {
		fmt.Println(color.InBlackOverYellow("UsingFiveM:"), GetUseFivem())
		router.HandleFunc("/api/{license}/{ipadresss}", API).Methods("GET")
	} else {
		fmt.Println(color.InBlackOverYellow("UsingFiveM:"), color.Green, GetUseFivem())
		router.HandleFunc("/api/{resourcename}/{license}/{ipadresss}", API).Methods("GET")
	}

	fmt.Println(color.InBlackOverYellow("Listening on port:"), color.Green+GetPort(), color.Reset)
	err := http.ListenAndServe(":"+GetPort(), router)
	if err != nil {
		fmt.Println(err)
	}
}
