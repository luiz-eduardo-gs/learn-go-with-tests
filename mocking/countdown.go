package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord       = "Go!"
	countdownStart  = 3
	countdownFinish = 1
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i >= countdownFinish; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{5 * time.Second, time.Sleep})
}
