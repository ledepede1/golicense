package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbURL string = GetDBUrl()

func CheckDB(license string, ipaddr string, actualip string, resourcename string) string {

	fmt.Println("Begninning license check...")

	if CheckIfLicenseExist(license, ipaddr) {
		fmt.Println("FIRST CHECK PASSED")
		if CheckIPListed(license, ipaddr, actualip, resourcename) {
			fmt.Println("SECOND CHECK PASSED")
			return "valid"
		} else {
			fmt.Println("CHECKS *NOT* PASSED")
			return "invalid"
		}
	}

	return "invalid"
}

func CheckIfLicenseExist(license string, ipaddr string) bool {
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var fetchedLicense string
	licenseQuerry := db.QueryRow("SELECT license FROM licenses WHERE license = ?", license).Scan(&fetchedLicense)

	if licenseQuerry == sql.ErrNoRows {
		fmt.Println("Invalid license: ", license)
		return false
	} else if licenseQuerry != nil {
		fmt.Println("Error: ", licenseQuerry)
		return false
	} else {
		fmt.Println("Successfully checked license:", license)
		return true
	}
}

func CheckIPListed(license string, sentipaddr string, actualip string, resourcename string) bool {
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var ipadress string
	ipcheckQuery := db.QueryRow("SELECT ipaddress FROM licenses WHERE license = ?", license).Scan(&ipadress)

	if actualip == ipadress {
		if ipcheckQuery == sql.ErrNoRows {
			fmt.Printf("No IP-Adress with that license%s", license)
			return false
		} else if sentipaddr == ipadress {
			fmt.Printf("IP: "+sentipaddr+" Setted up with: %s\n", license)

			if !GetUseFivem() {
				return true
			} else {
				if CheckFiveMResourceName(resourcename, license) {
					return true
				} else {
					fmt.Println("Resourcename: " + resourcename + " Is not setted up to license: " + license)
				}
			}
		}
	}

	return false
}

func CheckFiveMResourceName(resourcename string, license string) bool {

	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var savedresourcename string
	resourcenameQuery := db.QueryRow("SELECT resourcename FROM licenses WHERE license = ?", license).Scan(&savedresourcename)

	if resourcenameQuery == sql.ErrNoRows {
		fmt.Printf("No Resourcename with license%s", license)
		return false
	} else if savedresourcename == resourcename {
		fmt.Printf("Resourcename: "+savedresourcename+" Is setted up to license: %s\n", license)
		return true
	}
	return false
}