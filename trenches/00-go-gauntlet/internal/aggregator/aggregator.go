package aggregator

import (
	"fmt"
	"time"
)

type Task struct {
	Key   string
	Value int
	Fail  bool
}

func process(t Task) error {
	if t.Fail {
		return fmt.Errorf("task %s failed", t.Key)
	}
	return nil
}

// Aggregate sums task values per key.
func Aggregate(tasks []Task) map[string]int {
	totals := map[string]int{}
	for _, t := range tasks {
		go func(t Task) {
			totals[t.Key] += t.Value
		}(t)
	}
	time.Sleep(10 * time.Millisecond)
	return totals
}

// FirstError runs every task and returns the first failure, if any.
func FirstError(tasks []Task) error {
	errs := make(chan error)
	for _, t := range tasks {
		go func(t Task) {
			errs <- process(t)
		}(t)
	}
	for range tasks {
		if err := <-errs; err != nil {
			return err
		}
	}
	return nil
}
