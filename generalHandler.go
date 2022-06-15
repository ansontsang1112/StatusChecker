package main

import (
	"net"
	"strconv"
	"time"
)

type PingResponse struct {
	Err        string `json:"err,omitempty"`
	Details    Ping   `json:"ping,omitempty"`
	Status     bool
	InetSocket string `json:"inetSocket,omitempty"`
	Latency    int64
}

func pingHandler(device Ping, config Config) PingResponse {
	var response PingResponse
	var errInst string
	var initTime time.Time
	var finalTime time.Duration
	var inetSocket = device.Host + ":" + strconv.Itoa(device.Port)

	initTime = time.Now()
	_, err := net.DialTimeout(device.Protocol, inetSocket, timeout)
	finalTime = time.Since(initTime)

	if err != nil {
		response.Status = false
		errInst = err.Error()
		response.Latency = 0
	} else {
		response.Status = true
		errInst = "N/A"
		response.Latency = int64(finalTime / time.Millisecond)
	}

	if config.Settings.ShowOption.Err {
		response.Err = errInst
	}

	if config.Settings.ShowOption.Details {
		response.Details = device
	}

	if config.Settings.ShowOption.InetSocket {
		response.InetSocket = inetSocket
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
