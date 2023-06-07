package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func fetchName(n *name) error {
	u := url.URL{
		Scheme: "https",
		Host:   "names.mcquay.me",
		Path:   "api/v0/",
	}

	// @TODO: This retry code is repeated, move to shared function
	var resp *http.Response
	var err error

	for i := 0; i < retryCount; i++ {
		resp, err = httpClient.Get(u.String())

		if err != nil {
			log.Println("Error name: ", err)
		} else {
			break
		}
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// @TODO: Check resp.Body before decode, since there is a request limit reached when load testing
	// API returns string "You have reached maximum request limit." instead of JSON
	return json.NewDecoder(resp.Body).Decode(&n)
}
