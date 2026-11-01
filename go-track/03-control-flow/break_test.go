package main

import (
	"testing"
	"time"
)

// Attack tests. Run them with `make break` once `make test` passes.

func TestBreakGradeBoundaries(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"},
		{80, "B"}, {79, "C"}, {70, "C"},
		{69, "D"}, {60, "D"}, {59, "F"},
		{0, "F"}, {-1, "invalid"}, {101, "invalid"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestBreakSumToNothing(t *testing.T) {
	for _, n := range []int{0, -1, -5, -1000} {
		if got := SumTo(n); got != 0 {
			t.Errorf("SumTo(%d) = %d, want 0", n, got)
		}
	}
}

func TestBreakReverseUnicode(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"", ""},
		{"héllo", "olléh"},
		{"Go语言", "言语oG"},
		{"I ♥ Go", "oG ♥ I"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.s); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestBreakIsPrimeEdges(t *testing.T) {
	for _, n := range []int{-7, -1, 0, 1} {
		if IsPrime(n) {
			t.Errorf("IsPrime(%d) = true, want false", n)
		}
	}
}

// The largest prime that fits in an int32. A real system might check
// numbers like this thousands of times a second; one check gets 1 second.
func TestBreakIsPrimeFast(t *testing.T) {
	const n = 2_147_483_647
	result := make(chan bool, 1)
	go func() { result <- IsPrime(n) }()
	select {
	case got := <-result:
		if !got {
			t.Errorf("IsPrime(%d) = false, want true", n)
		}
	case <-time.After(time.Second):
		t.Fatalf("IsPrime(%d) took more than 1 second", n)
	}
}
