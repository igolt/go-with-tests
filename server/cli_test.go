package poker_test

import (
	"strings"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
	poker "github.com/igolt/go-with-tests/server"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)

		cli.PlayPoker()

		if len(playerStore.WinCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		asserts.AssertEqual(t, playerStore.WinCalls[0], "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)

		cli.PlayPoker()

		if len(playerStore.WinCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		asserts.AssertEqual(t, playerStore.WinCalls[0], "Cleo")
	})
}
