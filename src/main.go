package main

import (
	funcs "Pede-licensesystem/src/Funcs"
	"fmt"
	"net/http"

	"github.com/TwiN/go-color"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	if !funcs.GetUseFivem() {
		fmt.Println(color.InBlackOverYellow("UsingFiveM:"), funcs.GetUseFivem())
		router.HandleFunc("/api/{license}/{ipadresss}", API).Methods("GET")
	} else {
		fmt.Println(color.InBlackOverYellow("UsingFiveM:"), color.Green, funcs.GetUseFivem())
		router.HandleFunc("/api/{resourcename}/{license}/{ipadresss}", API).Methods("GET")
	}

	fmt.Println(color.InBlackOverYellow("Listening on port:"), color.Green+funcs.GetPort(), color.Reset)
	err := http.ListenAndServe(":"+funcs.GetPort(), router)
	if err != nil {
		fmt.Println(err)
	}
}
