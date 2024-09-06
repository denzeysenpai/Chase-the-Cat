package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not supported", http.StatusBadRequest)
	}
}
func handler2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not supported", http.StatusBadRequest)
	}
}
func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/home", handler)
	http.HandleFunc("/about", handler2)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
