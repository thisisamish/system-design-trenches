package aggregator

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func makeTasks(n, keys int) ([]Task, map[string]int) {
	tasks := make([]Task, n)
	want := map[string]int{}
	for i := range tasks {
		k := fmt.Sprintf("key-%d", i%keys)
		tasks[i] = Task{Key: k, Value: i}
		want[k] += i
	}
	return tasks, want
}

func TestAggregateTotals(t *testing.T) {
	tasks, want := makeTasks(10_000, 50)
	got := Aggregate(tasks)
	for k, w := range want {
		if got[k] != w {
			t.Fatalf("%s: got %d, want %d", k, got[k], w)
		}
	}
}

func TestAggregateHammered(t *testing.T) {
	tasks, want := makeTasks(1_000, 10)
	var wg sync.WaitGroup
	for c := 0; c < 100; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got := Aggregate(tasks)
			for k, w := range want {
				if got[k] != w {
					t.Errorf("%s: got %d, want %d", k, got[k], w)
					return
				}
			}
		}()
	}
	wg.Wait()
}

func TestFirstErrorReportsFailure(t *testing.T) {
	tasks, _ := makeTasks(100, 5)
	tasks[42].Fail = true
	if err := FirstError(tasks); err == nil {
		t.Fatal("expected an error")
	}
}

func TestFirstErrorLeaksNothing(t *testing.T) {
	before := runtime.NumGoroutine()
	tasks, _ := makeTasks(1_000, 5)
	for i := 0; i < len(tasks); i += 2 {
		tasks[i].Fail = true
	}
	for i := 0; i < 10; i++ {
		_ = FirstError(tasks)
	}
	time.Sleep(200 * time.Millisecond)
	if leaked := runtime.NumGoroutine() - before; leaked > 5 {
		t.Fatalf("%d goroutines still alive after FirstError returned", leaked)
	}
}
