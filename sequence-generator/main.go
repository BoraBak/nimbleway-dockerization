package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	count          = 1
	lock           sync.Mutex
	sequenceLength int
)

func setSequenceLengthFromEnvVar() {
	sequenceLengthStr := os.Getenv("SEQUENCE_LENGTH")
	if sequenceLengthStr == "" {
		sequenceLength = 5
	} else {
		var err error
		sequenceLength, err = strconv.Atoi(sequenceLengthStr)
		if err != nil {
			log.Fatal("Invalid sequence length:", err)
		}
	}
}

func handlerSequenceGenerator(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	w.Header().Set("X-Sequential-Number", strconv.Itoa(count))

	count++

	isResetCountWhenReachedSeqLen := count > sequenceLength
	if isResetCountWhenReachedSeqLen {
		count = 1
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	setSequenceLengthFromEnvVar()

	http.HandleFunc("/", handlerSequenceGenerator)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
