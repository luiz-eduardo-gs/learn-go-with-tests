package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		sleeperOps := &SpyCountdownOperations{}

		Countdown(sleeperOps, sleeperOps)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, sleeperOps.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleeperOps.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spy := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spy.Sleep}
	sleeper.Sleep()

	if spy.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spy.durationSlept)
	}
}
