package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// Returns new instance of Hello handler with given Logger.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
	}
	_, err = fmt.Fprintf(w, "hello %s\n", d)
	if err != nil {
		log.Println(err)
	}
}
