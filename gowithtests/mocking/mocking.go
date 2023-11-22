package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
	Write()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func (d *DefaultSleeper) Write(out io.Writer, i int) {
	if i >= 1 {
		fmt.Fprintln(out, i)
	} else {
		fmt.Fprint(out, "Go!")
	}
}
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i >= 0; i-- {
		sleeper.Write()
		sleeper.Sleep()
	}
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
