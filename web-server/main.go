package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	serverName  string
	sequenceNum int
)

func setServerNameFromEnvVar() {
	serverName = os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "web_server"
	}
}

func getSequentialNumber() int {
	resp, err := http.Get("http://sequence_generator:9090/")
	if err != nil {
		log.Fatalf("Error fetching sequential number: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	seqStr := resp.Header.Get("X-Sequential-Number")
	seq, err := strconv.Atoi(seqStr)
	if err != nil {
		log.Fatalf("Error converting sequential number to integer: %v", err)
	}

	return seq
}

func handlerWebServer(w http.ResponseWriter, r *http.Request) {
	if isRequestForMainPage(r) {
		return
	}

	sequenceNum = getSequentialNumber()
	currServerName := fmt.Sprintf("%s_%d", serverName, sequenceNum)
	fmt.Fprintf(w, "Server Name: %s\nSequential Number: %d\n", currServerName, sequenceNum)
}

func isRequestForMainPage(r *http.Request) bool {
	return r.URL.Path != "/"
}

func main() {
	setServerNameFromEnvVar()

	http.HandleFunc("/", handlerWebServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
