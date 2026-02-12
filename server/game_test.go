package poker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/igolt/go-with-tests/asserts"
	poker "github.com/igolt/go-with-tests/server"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
)

func TestGame_Start(t *testing.T) {
	dummyPlayerStore := &poker.StubPlayerStore{}

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		game.Start(5)

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)

	asserts.AssertEqual(t, len(store.WinCalls), 1)
	asserts.AssertEqual(t, store.WinCalls[0], winner)
}

func checkSchedulingCases(
	t *testing.T,
	cases []poker.ScheduledAlert,
	blindAlerter *poker.SpyBlindAlerter,
) {
	t.Helper()
	for i, expected := range cases {
		t.Run(fmt.Sprint(expected), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, &got, &expected)
		})
	}
}

func assertScheduledAlert(t testing.TB, got, expected *poker.ScheduledAlert) {
	t.Helper()

	if got.Amount != expected.Amount {
		t.Errorf("got amount %d, expected %d", got.Amount, expected.Amount)
	}

	if got.At != expected.At {
		t.Errorf("got scheduled time of %v, expected %v", got.At, expected.At)
	}
}
