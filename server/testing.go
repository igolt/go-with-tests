package poker

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
