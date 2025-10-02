package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
)

type stubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *stubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *stubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *stubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	server := NewPlayerServer(&stubPlayerStore{map[string]int{
		"Pepper": 20, "Floyd": 10,
	}, nil, nil})

	t.Run("returns Pepper's score", func(t *testing.T) {
		response := _GETPlayerScore(server, "Pepper")
		score := response.Body.String()

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		asserts.AssertEqual(t, score, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		response := _GETPlayerScore(server, "Floyd")
		score := response.Body.String()

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		asserts.AssertEqual(t, score, "10")
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		response := _GETPlayerScore(server, "NotFound")

		asserts.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}

func TestRecordWins(t *testing.T) {
	store := &stubPlayerStore{map[string]int{}, nil, nil}
	server := NewPlayerServer(store)

	t.Run("it record wins when POST", func(t *testing.T) {
		request := newRecordWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		asserts.AssertEqual(t, response.Code, http.StatusAccepted)
		asserts.AssertEqual(t, len(store.winCalls), 1)
		asserts.AssertEqual(t, store.winCalls[0], "Pepper")
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		expectedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := &stubPlayerStore{nil, nil, expectedLeague}
		server := NewPlayerServer(store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		asserts.AssertEqual(t, response.Code, http.StatusOK)
		assertLeague(t, got, expectedLeague)
		assertContentType(t, response, "application/json")
	})
}
