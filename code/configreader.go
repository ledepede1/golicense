package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Fivem bool   `json:"UsingFiveM"`
	DBUrl string `json:"DBUrl"`
	Port  string `json:"Port"`
}

func GetUseFivem() bool {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var cfgData Config

	json.Unmarshal(byteValue, &cfgData)

	return cfgData.Fivem
}

func GetDBUrl() string {

	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var cfgData Config

	json.Unmarshal(byteValue, &cfgData)

	return cfgData.DBUrl
}
func GetPort() string {

	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var cfgData Config

	json.Unmarshal(byteValue, &cfgData)

	return cfgData.Port
}
