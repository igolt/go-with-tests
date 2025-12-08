package main

import (
	"log"
	"net/http"

	poker "github.com/igolt/go-with-tests/server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	handler := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
