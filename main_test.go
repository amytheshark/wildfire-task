package main

import (
	"strings"
	"testing"
)

// @TODO: Update test to mock out http get request so we're not hitting real APIs
func TestFetchName(t *testing.T) {
	n := name{}
	fetchName(&n)

	if n.FirstName == "" {
		t.Errorf("Expected first name to be a non-empty string")
	}

	if n.LastName == "" {
		t.Errorf("Expected last name to be a non-empty string")
	}
}

// @TODO: Update test to mock out http get request so we're not hitting real APIs
func TestFetchJoke(t *testing.T) {
	n := name{
		"Tester",
		"Testerly",
	}

	j := joke{}
	fetchJoke(n, &j)

	if j.Type != "success" {
		t.Errorf("Expected joke API to return success response value, but got '%s'", j.Type)
	}

	if j.Value.Joke == "" {
		t.Errorf("Expected joke to be a string, but got no value")
	}

	// Originally written as separate tests, combined after I received a response with only first name
	if !strings.Contains(j.Value.Joke, n.FirstName) || !strings.Contains(j.Value.Joke, n.LastName) {
		t.Errorf(
			"Expected joke to contain first name '%s' and/or last name '%s', but got '%s'",
			n.FirstName,
			n.LastName,
			j.Value.Joke,
		)
	}
}
