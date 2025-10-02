package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
)

func TestRecordWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		asserts.AssertEqual(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		got := getLeagueFromResponse(t, response.Body)

		asserts.AssertEqual(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		assertLeague(t, got, []Player{{"Pepper", 3}})
	})
}
