package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func server(port int, devices Pings) {
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

		var displayResultSet, _ = json.Marshal(currentResultSetGenerator(devices))

		fmt.Fprint(w, string(displayResultSet))
	})

	// Server Handler
	fmt.Println("Try to start server at port " + strconv.Itoa(port))

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func currentResultSetGenerator(pings Pings) map[string]bool {
	var currentResultSet = make(map[string]bool)

	for i := 0; i < len(pings.Pings); i++ {
		currentResultSet[pings.Pings[i].Host], _ = pingHandler(pings.Pings[i])
	}

	return currentResultSet
}
