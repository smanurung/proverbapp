package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.Lshortfile)

	bibel := Bibel{}

	// List of endpoints.
	http.HandleFunc("/proverbs", bibel.HandleProverbs)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("listening to port 8808...")
	http.ListenAndServe(":8808", nil)
}
