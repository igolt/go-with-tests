package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := &Counter{}

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, expectedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, expected int) {
	t.Helper()
	if counter.Value() != expected {
		t.Errorf("got %d, expected %d", counter.Value(), expected)
	}
}
