package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spySleeper.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
	}
}
