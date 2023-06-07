package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type joke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		Id         int      `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}

func fetchJoke(n name, j *joke) error {
	u := url.URL{
		Scheme: "http",
		Host:   "joke.loc8u.com:8888",
		Path:   "joke",
	}

	q := u.Query()
	q.Set("limitTo", "nerdy")
	q.Set("firstName", n.FirstName)
	q.Set("lastName", n.LastName)
	u.RawQuery = q.Encode()

	// @TODO: This retry code is repeated, move to shared function
	var resp *http.Response
	var err error

	for i := 0; i < retryCount; i++ {
		resp, err = httpClient.Get(u.String())

		if err != nil {
			log.Println("Error joke: ", err)
		} else {
			break
		}
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&j)
}
