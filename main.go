package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello microservices!")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error in '/'", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Body: %s\n", body)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye microservices!")
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err.Error())
	}
}
