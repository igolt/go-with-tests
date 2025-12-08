package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	scanner     *bufio.Scanner
}

func NewCLI(playerStore PlayerStore, in io.Reader) *CLI {
	return &CLI{playerStore, bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin(extractWinner(cli.readLine()))
}

func (cli *CLI) readLine() string {
	cli.scanner.Scan()
	return cli.scanner.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
