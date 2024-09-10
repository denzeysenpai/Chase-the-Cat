package main

import (
	"log"
	"net/http"
)

type Car interface {
}

func main() {
	var arr [2]Car
	_ = arr
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
