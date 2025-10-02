package server

import "sync"

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[player]
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[player]++
}
