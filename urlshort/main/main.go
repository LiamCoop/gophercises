package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/liamcoop/gophercises/urlshort"
)

func main() {

	filePtr := flag.String("f", "", "file for input")
	flag.Parse()

	data, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	ext := strings.ToLower(strings.Split(*filePtr, ".")[1])
	var handler http.HandlerFunc

	if ext == "yaml" {
		handler, err = urlshort.YAMLHandler([]byte(data), mapHandler)
	} else if ext == "json" {
		handler, err = urlshort.JSONHandler([]byte(data), mapHandler)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
