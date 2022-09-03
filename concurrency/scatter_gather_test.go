package concurrency

import (
	"testing"
)

func Test_concurrencyLimiter_schedule(t *testing.T) {
	// limiter := New[string, int](2)

	// var limiter concurrencyLimiter = nil
	// work := func(arg int) func() string {
	// 	return func() string {
	// 		for i := 0; i < 5; i++ {
	// 			t.Logf("worker %d", arg)
	// 			if arg%2 == 0 {
	// 				time.Sleep(1 * time.Second)
	// 			} else {
	// 				time.Sleep(3 * time.Second)
	// 			}
	// 		}

	// 	}
	// }

	// for i := 0; i < 20; i++ {
	// 	limiter.schedule(i, work(i))
	// }

	// time.Sleep(1 * time.Hour)
}

func Test_doWork(t *testing.T) {
	coordinator()
}
