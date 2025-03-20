package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new test"

		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Delete("test")
		assertNoError(t, err)

		_, err = dictionary.Search("test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("test")

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertError(t testing.TB, err, expected error) {
	t.Helper()

	if err != expected {
		t.Errorf("got error %q expected %q", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("didn't expected to get an error. Got %q.", err)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, expected string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, expected)
}
