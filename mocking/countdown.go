package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownFinalWord = "Go!"
	countdownStart     = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, countdownFinalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
