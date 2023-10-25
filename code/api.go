package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func API(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	ipaddr := vars["ipadresss"]
	license := vars["license"]
	resourcename := vars["resourcename"]

	fmt.Println("\nIP Connected:", checkSentIP(request))

	if !GetUseFivem() {
		checkDatabase := CheckDB(string(license), string(ipaddr), checkSentIP(request), string(resourcename))
		fmt.Fprint(response, checkDatabase)
	} else {
		checkDatabase := CheckDB(string(license), string(ipaddr), checkSentIP(request), string(resourcename))
		fmt.Fprint(response, checkDatabase)
	}

}
