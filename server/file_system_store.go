package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	league      League
	database    io.WriteSeeker
	jsonEncoder json.Encoder
}

func NewFileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		db.Close()
		return nil, nil, err
	}
	return store, func() { db.Close() }, nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	if err := initializePlayerFile(file); err != nil {
		return nil, err
	}

	league, err := NewLeagueFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file system store: %v", err)
	}
	return &FileSystemPlayerStore{league, file, *json.NewEncoder(&tape{file})}, nil
}

func initializePlayerFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf(
			"failed to get file info while creating file system player store: %v",
			err,
		)
	}

	if stat.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := f.league.Find(playerName)

	if player == nil {
		return 0
	}
	return player.Wins
}

func (f *FileSystemPlayerStore) RecordWin(playerName string) {
	player := f.league.Find(playerName)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: playerName, Wins: 1})
	}
	f.jsonEncoder.Encode(f.league)
}
