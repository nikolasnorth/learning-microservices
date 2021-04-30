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

var port = "9090"

func main() {
	l := log.New(os.Stdout, "products: ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":" + port,
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
	l.Printf("listening on port %s...", port)

	// Listen for SIGINT
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until SIGINT happens
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
