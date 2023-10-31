package funcs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbURL string = GetDBUrl()

func CheckIfLicenseExist(license string) bool {
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var fetchedLicense string
	licenseQuerry := db.QueryRow("SELECT license FROM licenses WHERE license = ?", license).Scan(&fetchedLicense)

	if licenseQuerry == sql.ErrNoRows {
		fmt.Println("Invalid license: " + " '" + license + "'")
		return false
	} else {
		fmt.Println("Successfully checked license:" + " '" + fetchedLicense + "'")
		return true
	}
}

func CheckIPListed(license string, actualip string, resourcename string) bool {
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var ipadress string
	db.QueryRow("SELECT ipaddress FROM licenses WHERE license = ?", license).Scan(&ipadress)

	if actualip != ipadress {
		fmt.Println("IP: " + actualip + " Tried to connect with license: " + license + " But that IP not setted up to that license!")
		return false
	} else {
		fmt.Println("IP: " + ipadress + " Setted up with:" + " '" + license + "'")

		if !GetUseFivem() {
			return true
		} else {
			if CheckFiveMResourceName(resourcename, license) {
				return true
			} else {
				fmt.Println("Resourcename: " + "'" + resourcename + "'" + " Is not setted up to license: " + "'" + license + "'")
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
		fmt.Println("No Resourcename with license" + " '" + license + "'")
		return false
	} else if savedresourcename == resourcename {
		fmt.Println("Resourcename: " + "'" + savedresourcename + "'" + " Is setted up to license:" + " '" + license + "'")
		return true
	}
	return false
}
