package main

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_Found(t *testing.T) {
	data = make(map[string]string)
	data["/"] = "fred"
	hasNotFound := false
	notFound = func(w http.ResponseWriter, r *http.Request) {
		hasNotFound = true
	}

	redirectTarget := "noCalled"
	redirect = func(w http.ResponseWriter, r *http.Request, url string, code int) {
		redirectTarget = url
	}

	handler(nil, &http.Request{URL: &url.URL{Path: "/"}})

	if redirectTarget != "fred" {
		t.Logf("Expected %s, but got %s", "fred", redirectTarget)
		t.Fail()
	}

	if hasNotFound {
		t.Fail()
	}
}

func Test_Not_Found(t *testing.T) {
	data = make(map[string]string)
	data["/"] = "fred"
	hasNotFound := false
	notFound = func(w http.ResponseWriter, r *http.Request) {
		hasNotFound = true
	}

	redirectTarget := "noCalled"
	redirect = func(w http.ResponseWriter, r *http.Request, url string, code int) {
		redirectTarget = url
	}

	handler(nil, &http.Request{URL: &url.URL{Path: "/Fred"}})

	if redirectTarget != "noCalled" {
		t.Logf("Expected %s, but got %s", "noCalled", redirectTarget)
		t.Fail()
	}

	if !hasNotFound {
		t.Fail()
	}
}
