package poker

import (
	"fmt"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.Scores[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.WinCalls = append(s.WinCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (sa *ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", sa.Amount, sa.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: duration, Amount: amount})
}
