package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

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

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t testing.TB, body io.Reader) League {
	t.Helper()

	var league League
	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

func assertLeague(t testing.TB, got, expected League) {
	t.Helper()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, expected string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != expected {
		t.Errorf(
			"response did not have content-type of %s got %v",
			expected,
			response.Result().Header,
		)
	}
}
