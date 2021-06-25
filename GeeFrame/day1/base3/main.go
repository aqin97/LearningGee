package main

import (
	"LearningGee/GeeFrame/day1/base3/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	log.Fatal(r.Run("127.0.0.1:8080"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
