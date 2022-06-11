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
	Id          string `json:"id,omitempty"`
	Host        string `json:"host,omitempty"`
	Port        int    `json:"port,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Description string `json:"description,omitempty"`
}

type Config struct {
	HttpConfig HttpConfig `json:"httpConfig"`
	Settings   Settings   `json:"settings"`
}

type Settings struct {
	ShowOption ShowOption `json:"show"`
}

type ShowOption struct {
	Err     bool `json:"err"`
	Details bool `json:"details"`
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
