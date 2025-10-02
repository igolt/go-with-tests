package server

import (
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

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	asserts.AssertEqual(t, response.Body.String(), "3")
}
