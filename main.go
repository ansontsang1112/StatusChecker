package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const timeout time.Duration = time.Duration(2 * time.Second)

func main() {
	var pingDevices = jsonHandler("pingDevices.json")

	for i := 0; i < len(pingDevices.Pings); i++ {
		fmt.Println(pingHandler(pingDevices.Pings[i]))
	}

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
