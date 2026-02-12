package poker

import "time"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter     BlindAlerter
	playerStore PlayerStore
}

func NewTexasHoldem(alerter BlindAlerter, playerStore PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter, playerStore}
}

func (p *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.playerStore.RecordWin(winner)
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
