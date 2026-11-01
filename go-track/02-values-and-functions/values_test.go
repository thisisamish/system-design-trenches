package main

import (
	"math"
	"testing"
)

// closeTo reports whether two floats are equal, give or take rounding error.
func closeTo(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		c, want float64
	}{
		{100, 212},
		{0, 32},
		{-40, -40},
		{37, 98.6},
		{21.5, 70.7},
	}
	for _, tt := range tests {
		if got := CelsiusToFahrenheit(tt.c); !closeTo(got, tt.want) {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.want)
		}
	}
}

func TestSplitCents(t *testing.T) {
	tests := []struct {
		total, dollars, cents int
	}{
		{1999, 19, 99},
		{500, 5, 0},
		{7, 0, 7},
		{0, 0, 0},
	}
	for _, tt := range tests {
		d, c := SplitCents(tt.total)
		if d != tt.dollars || c != tt.cents {
			t.Errorf("SplitCents(%d) = %d, %d, want %d, %d", tt.total, d, c, tt.dollars, tt.cents)
		}
	}
}

func TestPercentOf(t *testing.T) {
	tests := []struct {
		part, whole int
		want        float64
	}{
		{50, 200, 25},
		{1, 3, 100.0 / 3},
		{3, 3, 100},
		{0, 9, 0},
		{7, 4, 175},
	}
	for _, tt := range tests {
		if got := PercentOf(tt.part, tt.whole); !closeTo(got, tt.want) {
			t.Errorf("PercentOf(%d, %d) = %v, want %v", tt.part, tt.whole, got, tt.want)
		}
	}
}

func TestMillisFromSeconds(t *testing.T) {
	tests := []struct {
		s    int32
		want int64
	}{
		{90, 90_000},
		{1, 1000},
		{0, 0},
		{3600, 3_600_000},
	}
	for _, tt := range tests {
		if got := MillisFromSeconds(tt.s); got != tt.want {
			t.Errorf("MillisFromSeconds(%d) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{10, 15, 12},
		{10, 20, 15},
		{0, 0, 0},
		{0, 1, 0},
		{7, 7, 7},
	}
	for _, tt := range tests {
		if got := Average(tt.a, tt.b); got != tt.want {
			t.Errorf("Average(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
