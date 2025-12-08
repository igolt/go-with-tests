package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/igolt/go-with-tests/server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	cli := poker.NewCLI(store, os.Stdin)

	fmt.Println("Let's play poker")
	fmt.Println("Type '{name} wins' to record a player win")

	cli.PlayPoker()
}
