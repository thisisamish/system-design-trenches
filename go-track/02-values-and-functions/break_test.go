package main

import (
	"math"
	"testing"
)

// Attack tests. Run them with `make break` once `make test` passes.

func TestBreakMillisFromSecondsLong(t *testing.T) {
	tests := []struct {
		s    int32
		want int64
	}{
		{86_400 * 30, 2_592_000_000},       // a 30-day timeout
		{3_000_000, 3_000_000_000},         // about 35 days
		{math.MaxInt32, 2_147_483_647_000}, // the largest int32 there is
	}
	for _, tt := range tests {
		if got := MillisFromSeconds(tt.s); got != tt.want {
			t.Errorf("MillisFromSeconds(%d) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestBreakAverageHuge(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{math.MaxInt, math.MaxInt, math.MaxInt},
		{math.MaxInt - 1, math.MaxInt - 3, math.MaxInt - 2},
		{math.MaxInt, 1, math.MaxInt/2 + 1},
		{1 << 62, 1 << 62, 1 << 62},
	}
	for _, tt := range tests {
		if got := Average(tt.a, tt.b); got != tt.want {
			t.Errorf("Average(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestBreakSplitCentsHuge(t *testing.T) {
	d, c := SplitCents(math.MaxInt)
	if d != math.MaxInt/100 || c != math.MaxInt%100 {
		t.Errorf("SplitCents(%d) = %d, %d, want %d, %d", math.MaxInt, d, c, math.MaxInt/100, math.MaxInt%100)
	}
}
