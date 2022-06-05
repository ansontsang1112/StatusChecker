package main

import (
	"net"
	"strconv"
)

type PingResponse struct {
	Ping   Ping
	Status bool
	Err    string
}

func pingHandler(device Ping) PingResponse {
	var response PingResponse
	var inetSocket = device.Host + ":" + strconv.Itoa(device.Port)
	_, err := net.DialTimeout(device.Protocol, inetSocket, timeout)

	response.Ping = device
	if err != nil {
		response.Status = false
		response.Err = err.Error()
	} else {
		response.Status = true
		response.Err = "N/A"
	}

	return response
}

func currentResultSetGenerator(pings Pings) map[string]PingResponse {
	var pingResponses = make(map[string]PingResponse)

	for i := 0; i < len(pings.Pings); i++ {
		pingResponses[pings.Pings[i].Id] = pingHandler(pings.Pings[i])
	}

	return pingResponses
}
