package main

import (
	"time"
)

const timeout = 1 * time.Second
const settingsLocation = "settings"

func main() {
	// Read Json Configuration
	var pingDevices = pingJsonHandler("pingDevices.json") // Init Devices
	var configurations = configJsonHandler("config.json") // Init Config

	server(pingDevices, configurations)
}
