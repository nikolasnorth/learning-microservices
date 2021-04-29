package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	l.Printf("listening on port %s...", PORT)

	s := &http.Server{
		Addr:         ":" + PORT,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Listen for SIGINT signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Will block until SIGINT happens
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// Allow 30s grace period before shutting down server. Accept no new connections during this time.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		l.Fatal(err)
	}
}
