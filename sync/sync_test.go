package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		assertCount(t, c, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		c := NewCounter()

		var wg sync.WaitGroup
		// quantas go routines esperar
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		// espera todas go routines acabarem
		wg.Wait()

		assertCount(t, c, wantedCount)
	})
}

func assertCount(t testing.TB, c *Counter, want int) {
	t.Helper()

	if c.Value() != want {
		t.Errorf("got %d, want %d", c.Value(), want)
	}
}
