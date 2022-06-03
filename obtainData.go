package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Pings struct {
	Pings []Ping `json:"devices"`
}

type Ping struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}

func jsonHandler(fileName string) Pings {
	var pings Pings
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fileName)
	defer jsonFile.Close() // Close Json File

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &pings)

	return pings
}
