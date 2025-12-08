package main

import (
	"log"
	"net/http"
	"os"

	"github.com/igolt/go-with-tests/server"
)

const dbFileName = "game.db.json"

func main() {
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatalf("failed opening %s %v", dbFileName, err)
	}

	store, err := server.NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatal(err)
	}

	handler := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
