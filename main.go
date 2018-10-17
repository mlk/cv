package main

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var data map[string]string

var notFound = http.NotFound
var redirect = http.Redirect
var printf = log.Printf

func handler(w http.ResponseWriter, r *http.Request) {
	url, ok := data[r.URL.Path]

	printf("%s", r.URL.Path)
	if !ok {
		notFound(w, r)
	} else {
		redirect(w, r, url, http.StatusFound)
	}
}

func main() {
	log.Print("Started")
	filename, _ := filepath.Abs("./cvs.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
