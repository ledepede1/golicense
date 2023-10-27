package main

import (
	funcs "Pede-licensesystem/src/Funcs"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func API(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	ipaddr := vars["ipadresss"]
	license := vars["license"]
	resourcename := vars["resourcename"]

	fmt.Println("\n----------------------------------------------------------------------------------------")
	fmt.Println("IP Connected:", funcs.CheckSentIP(request))

	if !funcs.GetUseFivem() {
		checkDatabase := funcs.CheckDB(string(license), string(ipaddr), funcs.CheckSentIP(request), string(resourcename))
		fmt.Fprint(response, checkDatabase)
	} else {
		checkDatabase := funcs.CheckDB(string(license), string(ipaddr), funcs.CheckSentIP(request), string(resourcename))
		fmt.Fprint(response, checkDatabase)
	}

}
