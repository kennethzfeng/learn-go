// HTTP Server

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	port := os.Getenv("PORT")
	log.Printf("Getting PORT %d", port)

	http.HandleFunc("/", Index)

	log.Printf("Listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
