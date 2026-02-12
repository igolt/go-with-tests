package poker_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
	poker "github.com/igolt/go-with-tests/server"
)

var dummyOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		gotPrompt := out.String()

		asserts.AssertEqual(t, gotPrompt, poker.PlayerPrompt)
	})

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, dummyOut, game)

		cli.PlayPoker()

		asserts.AssertEqual(t, game.StartedWith, 5)
		asserts.AssertEqual(t, game.FinishedWith, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, dummyOut, game)

		cli.PlayPoker()

		asserts.AssertEqual(t, game.StartedWith, 5)
		asserts.AssertEqual(t, game.FinishedWith, "Cleo")
	})

	t.Run(
		"it prints an error when a non numeric value is entered and does not start the game",
		func(t *testing.T) {
			out := &bytes.Buffer{}
			in := strings.NewReader("Pies")
			game := &poker.GameSpy{}

			cli := poker.NewCLI(in, out, game)
			cli.PlayPoker()

			if game.StartCalled {
				t.Errorf("game should not have started")
			}

			gotPrompt := out.String()
			asserts.AssertEqual(t, gotPrompt, poker.PlayerPrompt+poker.BadPlayerInputErrMsg)
		},
	)
}
