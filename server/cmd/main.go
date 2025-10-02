package main

import (
	"log"
	"net/http"

	"github.com/igolt/go-with-tests/server"
)

func main() {
	handler := server.NewPlayerServer(&server.InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", handler))
}
