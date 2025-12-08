package poker

import (
	"os"
	"testing"

	"github.com/igolt/go-with-tests/asserts"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer removeFile()

		store := newFileSystemStore(t, file)

		expected := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := store.GetLeague()

		assertLeague(t, got, expected)
	})

	t.Run("get player score", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer removeFile()

		store := newFileSystemStore(t, file)

		got := store.GetPlayerScore("Chris")

		asserts.AssertEqual(t, got, 33)
	})

	t.Run("store wins for existing player", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer removeFile()
		player := "Chris"

		store := newFileSystemStore(t, file)

		store.RecordWin(player)

		got := store.GetPlayerScore(player)

		asserts.AssertEqual(t, got, 34)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer removeFile()

		store := newFileSystemStore(t, file)
		player := "Pepper"

		store.RecordWin(player)

		got := store.GetPlayerScore(player)

		asserts.AssertEqual(t, got, 1)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		file, removeFile := createTempFile(t, "")
		defer removeFile()

		_, err := NewFileSystemPlayerStore(file)
		if err != nil {
			t.Fatalf(
				"didn't expect NewFileSystemPlayerStore to return an error when creating with an empty file but got: %v",
				err,
			)
		}
	})

	t.Run("league sorted in decreasing order of wins", func(t *testing.T) {
		file, removeFile := createTempFile(
			t,
			`[{"Name": "Cleo","Wins":10},{"Name":"Chris","Wins":33}]`,
		)
		defer removeFile()

		store := newFileSystemStore(t, file)

		got := store.GetLeague()

		expected := League{{"Chris", 33}, {"Cleo", 10}}

		assertLeague(t, got, expected)
	})
}

func newFileSystemStore(t testing.TB, file *os.File) *FileSystemPlayerStore {
	t.Helper()

	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		t.Fatalf("failed to create file system store: %v", err)
	}
	return store
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	file, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}

	file.Write([]byte(initialData))
	file.Seek(0, 0)

	removeFile := func() {
		file.Close()
		os.Remove(file.Name())
	}

	return file, removeFile
}
