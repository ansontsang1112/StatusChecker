package main

import (
	"net"
	"strconv"
)

type PingResponse struct {
	Err        string `json:"err,omitempty"`
	Details    Ping   `json:"ping,omitempty"`
	Status     bool
	InetSocket string
}

func pingHandler(device Ping, config Config) PingResponse {
	var response PingResponse
	var errInst string
	var inetSocket = device.Host + ":" + strconv.Itoa(device.Port)

	_, err := net.DialTimeout(device.Protocol, inetSocket, timeout)

	response.InetSocket = inetSocket

	if err != nil {
		response.Status = false
		errInst = err.Error()
	} else {
		response.Status = true
		errInst = "N/A"
	}

	if config.Settings.ShowOption.Err {
		response.Err = errInst
	}

	if config.Settings.ShowOption.Details {
		response.Details = device
	}

	return response
}

func currentResultSetGenerator(pings Pings, config Config) map[string]PingResponse {
	var pingResponses = make(map[string]PingResponse)

	for i := 0; i < len(pings.Pings); i++ {
		pingResponses[pings.Pings[i].Id] = pingHandler(pings.Pings[i], config)
	}

	return pingResponses
}
