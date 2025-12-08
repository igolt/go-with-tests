package server

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeagueFromReader(reader io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("failed to parse league: %v", err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for idx := range l {
		p := &l[idx]
		if p.Name == name {
			return p
		}
	}
	return nil
}
