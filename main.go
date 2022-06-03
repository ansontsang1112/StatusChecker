package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

const timeout time.Duration = time.Duration(2 * time.Second)

type Pings struct {
	Pings []Ping `json:"devices"`
}

type Ping struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}

func main() {
	var pingDevices = jsonHandler("pingDevices.json")

	for i := 0; i < len(pingDevices.Pings); i++ {
		fmt.Println(pingHandler(pingDevices.Pings[i]))
	}

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

func pingHandler(device Ping) (bool, string) {
	var inetSocket = device.Host + ":" + strconv.Itoa(device.Port)
	_, err := net.DialTimeout(device.Protocol, inetSocket, timeout)

	if err != nil {
		return false, err.Error()
	} else {
		return true, ""
	}
}
