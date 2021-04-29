package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nikolasnorth/microservices/handlers"
)

var PORT = "9090"

func main() {
	l := log.New(os.Stdout, "hello-handler:", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	log.Printf("listening on port %s...", PORT)
	err := http.ListenAndServe(":"+PORT, sm)
	if err != nil {
		log.Fatal(err)
	}
}
