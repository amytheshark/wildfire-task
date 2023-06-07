package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}
var retryCount = 5

func main() {
	http.HandleFunc("/", completeTask)
	http.ListenAndServe(":5000", nil)
}

func completeTask(w http.ResponseWriter, r *http.Request) {
	n := name{}
	err := fetchName(&n)

	if err != nil {
		log.Println("Error name2:", err)
		io.WriteString(w, "Error: Encountered issue fetching name, please try again later.")
		return
	}

	j := joke{}
	errJoke := fetchJoke(n, &j)

	if errJoke != nil {
		log.Println("Error joke2:", errJoke)
		io.WriteString(w, "Error: Encountered issue fetching joke, please try again later.")
		return
	}

	io.WriteString(w, j.Value.Joke)
}
