package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var filename string = "config.json"

type Config struct {
	Fivem bool   `json:"UsingFiveM"`
	DBUrl string `json:"DBUrl"`
	Port  string `json:"Port"`
}

func GetCfg(returnValue string) (string, bool) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var cfgData Config

	json.Unmarshal(byteValue, &cfgData)

	switch returnValue {
	case "Fivem":
		return "", cfgData.Fivem
	case "DBUrl":
		return cfgData.DBUrl, false
	case "Port":
		return cfgData.Port, false
	default:
		return "Error", false
	}
}

func GetUseFivem() bool {
	_, getCfg := GetCfg("Fivem")
	return getCfg
}

func GetDBUrl() string {
	getCfg, _ := GetCfg("DBUrl")
	return getCfg
}

func GetPort() string {
	getCfg, _ := GetCfg("Port")
	return getCfg
}
