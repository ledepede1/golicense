package main

import (
	funcs "Pede-licensesystem/src/Funcs"
	"fmt"
	"net/http"

	"github.com/TwiN/go-color"
	"github.com/gorilla/mux"
)

func main() {
	var urlVars string = "/api/{license}"
	var port string = funcs.GetPort()
	var usingFivem = funcs.GetUseFivem()

	if usingFivem {
		urlVars = "/api/{resourcename}/{license}"
	}

	router := mux.NewRouter()

	fmt.Println(color.InBlackOverYellow("UsingFiveM:"), usingFivem)
	router.HandleFunc(urlVars, API).Methods("GET")

	fmt.Println(color.InBlackOverYellow("Listening on port:"), color.Green+port, color.Reset)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
