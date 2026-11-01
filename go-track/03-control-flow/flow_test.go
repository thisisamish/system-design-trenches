package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{30, "F"},
		{150, "invalid"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestSumTo(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{1, 1},
		{4, 10},
		{10, 55},
		{100, 5050},
	}
	for _, tt := range tests {
		if got := SumTo(tt.n); got != tt.want {
			t.Errorf("SumTo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"hello", "olleh"},
		{"Go", "oG"},
		{"racecar", "racecar"},
		{"a b c", "c b a"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.s); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{17, true},
		{25, false},
		{97, true},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
