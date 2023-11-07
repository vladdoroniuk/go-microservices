package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vladdoroniuk/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	mux := http.NewServeMux()
	mux.Handle("/hello", hh)
	mux.Handle("/goodbye", gh)

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err.Error())
	}
}
