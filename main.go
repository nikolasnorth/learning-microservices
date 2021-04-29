package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var PORT = "9090"

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/goodbye", handleGoodbye)

	log.Printf("listening on port %s...", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	log.Println("goodbye world")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("hello world")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
	}
	_, err = fmt.Fprintf(w, "hello %s", d)
	if err != nil {
		log.Println(err)
	}
}
