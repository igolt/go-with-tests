package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
)

type stubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *stubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *stubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func TestGETPlayers(t *testing.T) {
	playerServer := NewPlayerServer(&stubPlayerStore{map[string]int{
		"Pepper": 20, "Floyd": 10,
	}, nil})

	t.Run("returns Pepper's score", func(t *testing.T) {
		response := _GETPlayerScore(playerServer, "Pepper")
		score := response.Body.String()

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		asserts.AssertEqual(t, score, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		response := _GETPlayerScore(playerServer, "Floyd")
		score := response.Body.String()

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		asserts.AssertEqual(t, score, "10")
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		response := _GETPlayerScore(playerServer, "NotFound")

		asserts.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}

func TestRecordWins(t *testing.T) {
	store := &stubPlayerStore{map[string]int{}, nil}
	playerServer := NewPlayerServer(store)

	t.Run("it record wins when POST", func(t *testing.T) {
		request := newRecordWinRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		asserts.AssertEqual(t, response.Code, http.StatusAccepted)
		asserts.AssertEqual(t, len(store.winCalls), 1)
		asserts.AssertEqual(t, store.winCalls[0], "Pepper")
	})
}

func _GETPlayerScore(playerServer *PlayerServer, player string) *httptest.ResponseRecorder {
	request := newGetScoreRequest(player)
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)
	return response
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newRecordWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}
