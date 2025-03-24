package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	buffer bytes.Buffer
	Calls  []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return s.buffer.Write(p)
}

func (s *SpyCountdownOperations) String() string {
	return s.buffer.String()
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const (
	write = "write"
	sleep = "sleep"
)

type SpyTime struct {
	durationSlept time.Duration
}

func TestCountdown(t *testing.T) {
	spySleeper := &SpyCountdownOperations{}

	Countdown(spySleeper, spySleeper)

	output := spySleeper.String()
	expectedOuput := `3
2
1
Go!`

	expectedCallOrder := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if output != expectedOuput {
		t.Errorf("expected output %q but got %q", expectedOuput, output)
	}

	if !reflect.DeepEqual(expectedCallOrder, spySleeper.Calls) {
		t.Errorf("expected calls %v but got %v", expectedCallOrder, spySleeper.Calls)
	}
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
