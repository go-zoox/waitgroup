package waitgroup

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	type test struct {
		name string
		size int
	}

	tests := []test{
		{"sing;e", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := New(test.size)
			if wg.PendingCount() != 0 {
				t.Errorf("PendingCount() should be 0, got %d", wg.PendingCount())
			}

			wg.Add(func() {
				t.Log("hello")
			})

			if wg.PendingCount() != 1 {
				t.Errorf("PendingCount() should be 1, got %d", wg.PendingCount())
			}

			wg.Wait()

			if wg.PendingCount() != 0 {
				t.Errorf("PendingCount() should be 0, got %d", wg.PendingCount())
			}
		})
	}
}

func TestWaitGroupRealWorld(t *testing.T) {
	wg := New(3)
	jobs := []func(){}

	for i := 0; i < 10; i++ {
		index := i
		jobs = append(jobs, func() {
			time.Sleep(time.Second)
			fmt.Println(index)
		})
	}

	wg.Add(jobs...)

	wg.Wait()
}
