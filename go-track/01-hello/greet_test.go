package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Gopher")
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("Greeting(%q) = %q, want %q", "Gopher", got, want)
	}

	got = Greeting("Ada")
	want = "Hello, Ada!"
	if got != want {
		t.Errorf("Greeting(%q) = %q, want %q", "Ada", got, want)
	}
}

func TestShout(t *testing.T) {
	got := Shout("go")
	want := "GO!!!"
	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "go", got, want)
	}

	got = Shout("Hello there")
	want = "HELLO THERE!!!"
	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "Hello there", got, want)
	}
}

func TestIntroduce(t *testing.T) {
	got := Introduce("Gopher", 13)
	want := "Gopher is 13 years old"
	if got != want {
		t.Errorf("Introduce(%q, %d) = %q, want %q", "Gopher", 13, got, want)
	}

	got = Introduce("Ada", 36)
	want = "Ada is 36 years old"
	if got != want {
		t.Errorf("Introduce(%q, %d) = %q, want %q", "Ada", 36, got, want)
	}
}
