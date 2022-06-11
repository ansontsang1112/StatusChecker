package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func server(devices Pings, config Config) {

	// All HTTP Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 Page Not Found", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var displayResultSet, _ = json.Marshal(currentResultSetGenerator(devices, config))

		fmt.Fprint(w, string(displayResultSet))
	})

	// Server Handler
	fmt.Println("Starting server at port " + strconv.Itoa(config.HttpConfig.Port))

	err := http.ListenAndServe(":"+strconv.Itoa(config.HttpConfig.Port), nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
