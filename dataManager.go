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
	Id          string `json:"id"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}

type Config struct {
	HttpConfig HttpConfig `json:"httpConfig"`
}

type HttpConfig struct {
	Port int `json:"port"`
}

func pingJsonHandler(fileName string) Pings {
	var pings Pings
	jsonFile, err := os.Open(settingsLocation + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close() // Close Json File

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &pings)

	return pings
}

func configJsonHandler(fileName string) Config {
	var config Config
	jsonFile, err := os.Open(settingsLocation + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)

	return config
}
