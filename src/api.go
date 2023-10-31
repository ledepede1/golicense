package main

import (
	funcs "Pede-licensesystem/src/Funcs"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func API(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	license := vars["license"]
	resourcename := vars["resourcename"]

	fmt.Println("\n----------------------------------------------------------------------------------------")
	fmt.Println("IP Connected:", funcs.CheckSentIP(request))

	checkDatabase := CheckDB(string(license), funcs.CheckSentIP(request), string(resourcename))
	fmt.Fprint(response, checkDatabase)
}

func CheckDB(license string, actualip string, resourcename string) string {

	fmt.Println("Begninning license check...")
	fmt.Println("")

	if funcs.CheckIfLicenseExist(license) {
		fmt.Println("FIRST CHECK PASSED")
		if funcs.CheckIPListed(license, actualip, resourcename) {
			fmt.Println("SECOND CHECK PASSED")
			fmt.Println("\nAll checks passed")
			fmt.Println("----------------------------------------------------------------------------------------")
			return "valid"
		} else {
			fmt.Println("\nCHECKS *NOT* PASSED")
			fmt.Println("----------------------------------------------------------------------------------------")
			return "invalid"
		}
	}
	fmt.Println("\nCHECKS *NOT* PASSED")
	fmt.Println("----------------------------------------------------------------------------------------")
	return "invalid"
}
