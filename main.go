package main

import (
	"net"
	"strconv"
	"time"
)

const timeout = 1 * time.Second

func main() {
	// Read Json Configuration
	var pingDevices = pingJsonHandler("pingDevices.json")
	var configurations = configJsonHandler("config.json")

	server(configurations.HttpConfig.Port, pingDevices)
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
