package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	scanner *bufio.Scanner
	out     io.Writer
	game    Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{bufio.NewScanner(in), out, game}
}

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
)

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	userInput := cli.readLine()
	winner := extractWinner(userInput)

	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.scanner.Scan()
	return cli.scanner.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
