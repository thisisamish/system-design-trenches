package logparser

import (
	"errors"
	"fmt"
	"testing"
)

// noPanic fails the test instead of letting a panic kill the whole binary.
func noPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()
	f()
}

func msgs(es []Entry) []string {
	out := make([]string, len(es))
	for i, e := range es {
		out[i] = e.Msg
	}
	return out
}

func ring(cap, pushes int) *RingBuffer {
	r := NewRingBuffer(cap)
	for i := 0; i < pushes; i++ {
		r.Push(Entry{Msg: fmt.Sprint(i)})
	}
	return r
}

func TestParseWellFormed(t *testing.T) {
	e, err := Parse("WARN|user=7,region=eu|disk almost full")
	if err != nil || e.Level != "WARN" || e.Fields["user"] != "7" || e.Msg != "disk almost full" {
		t.Fatalf("got %+v, %v", e, err)
	}
}

func TestParseNoFieldsThenSet(t *testing.T) {
	noPanic(t, func() {
		e, err := Parse("INFO||hello")
		if err != nil {
			t.Fatal(err)
		}
		e.Set("trace", "abc")
		if e.Fields["trace"] != "abc" {
			t.Fatal("Set lost the value")
		}
	})
}

func TestParseMalformed(t *testing.T) {
	for _, line := range []string{"garbage", "INFO|only-two", "INFO|novalue|msg", ""} {
		noPanic(t, func() {
			if _, err := Parse(line); !errors.Is(err, ErrMalformed) {
				t.Fatalf("Parse(%q) err = %v, want ErrMalformed", line, err)
			}
		})
	}
}

func TestParseAllReportsBadLine(t *testing.T) {
	noPanic(t, func() {
		es, err := ParseAll([]string{"INFO||ok", "nope", "WARN||ok2"})
		if !errors.Is(err, ErrMalformed) {
			t.Fatalf("err = %v, want it to wrap ErrMalformed", err)
		}
		for _, e := range es {
			if e == nil {
				t.Fatal("nil entry in result")
			}
		}
	})
}

func TestPointersAreDistinct(t *testing.T) {
	ps := Pointers([]Entry{{Msg: "a"}, {Msg: "b"}, {Msg: "c"}})
	got := fmt.Sprint(ps[0].Msg, ps[1].Msg, ps[2].Msg)
	if got != "a b c" && got != "abc" {
		t.Fatalf("got %q, want entries a, b, c", got)
	}
}

func TestRingSnapshotOrderAfterWrap(t *testing.T) {
	got := msgs(ring(3, 5).Snapshot())
	want := []string{"2", "3", "4"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRingSnapshotIsACopy(t *testing.T) {
	r := ring(3, 3)
	s := r.Snapshot()
	s[0].Msg = "tampered"
	if r.Snapshot()[0].Msg == "tampered" {
		t.Fatal("caller mutation leaked into the buffer")
	}
}

func TestRingLastAfterWrap(t *testing.T) {
	got := msgs(ring(3, 5).Last(2))
	want := []string{"3", "4"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRingLastMoreThanHeld(t *testing.T) {
	noPanic(t, func() {
		if n := len(ring(3, 2).Last(10)); n != 2 {
			t.Fatalf("got %d entries, want 2", n)
		}
	})
}

func TestRingZeroCapacity(t *testing.T) {
	noPanic(t, func() {
		r := NewRingBuffer(0)
		r.Push(Entry{Msg: "x"})
		if n := len(r.Snapshot()); n != 0 {
			t.Fatalf("got %d entries, want 0", n)
		}
	})
}
