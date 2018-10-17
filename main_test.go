package main

import (
	"net/http"
	"net/url"
	"testing"
)

var hasNotFound = false
var redirectTarget = "noCalled"

func init() {
	data = make(map[string]string)
	data["/"] = "fred"

	notFound = func(w http.ResponseWriter, r *http.Request) {
		hasNotFound = true
	}

	redirect = func(w http.ResponseWriter, r *http.Request, url string, code int) {
		redirectTarget = url
	}
}

func resetFake() {
	hasNotFound = false
	redirectTarget = "noCalled"
}

func assertEquals(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Logf("Expected %s, but got %s", expected, actual)
		t.Fail()
	}
}

func Test_Found(t *testing.T) {
	resetFake()

	handler(nil, &http.Request{URL: &url.URL{Path: "/"}})

	assertEquals(t, "fred", redirectTarget)

	if hasNotFound {
		t.Fail()
	}
}

func Test_Not_Found(t *testing.T) {
	resetFake()

	handler(nil, &http.Request{URL: &url.URL{Path: "/Fred"}})

	assertEquals(t, "noCalled", redirectTarget)

	if !hasNotFound {
		t.Fail()
	}
}
