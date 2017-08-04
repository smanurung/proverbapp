package main

import (
	"log"

	b "github.com/sonnythehottest/bibel"
)

func main() {
	log.SetFlags(log.Lshortfile)

	bibel := b.Bibel{}

	// List of endpoints.
	bibel.HandleProverbs()
}
