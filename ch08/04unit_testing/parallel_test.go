package main

import (
	"testing"
	"time"
)

func Parallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func Parallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func Parallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallel(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip-yade")
	}
	Parallel_1(t)
	Parallel_2(t)
	Parallel_3(t)
}
