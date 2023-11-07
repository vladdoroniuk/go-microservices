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

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello microservices!")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in '/hello'", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s\n", body)
}
