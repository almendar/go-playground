package concurrency

import (
	"context"
	"fmt"
	"time"
)

type concurrencyLimiter[FROM, TO any] struct {
	max int
	ch  chan bool
}

func New[FROM, TO any](max int) concurrencyLimiter[FROM, TO] {
	return concurrencyLimiter[FROM, TO]{
		max,
		make(chan bool, max),
	}
}

// func (it *concurrencyLimiter[FROM, TO]) schedule(arg FROM,
// 	action func(from FROM) TO) chan TO {
// 	retCh := make(chan TO, 1)
// 	go func(rcv chan TO) {
// 		it.ch <- true
// 		rcv <- action(arg)
// 		<-it.ch
// 	}(retCh)
// }

type workResult struct {
	done int
	err  error
}

func doWork(ctx context.Context, arg string) workResult {
	fmt.Printf("doWork: %s\n", arg)
	resultCh := make(chan workResult, 1)

	go func() {
		fmt.Printf("doWork:%s sleep-start \n", arg)
		if len(arg)%2 == 0 {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(6 * time.Second)
		}
		fmt.Printf("doWork:%s sleep-end \n", arg)
		resultCh <- workResult{len(arg), nil}
		fmt.Printf("doWork:%s result send to ch \n", arg)
		close(resultCh)
	}()

	select {
	case computedVal := <-resultCh:
		return computedVal
	case <-ctx.Done():
		return workResult{0, fmt.Errorf("timeout: %v", arg)}
	}
}

func coordinator() {
	listOfArgs := []string{"a", "ab", "abc", "abcd"}
	numberOfItems := len(listOfArgs)
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	results := make(chan workResult, numberOfItems)

	for _, v := range listOfArgs {
		fmt.Printf("Processing %s\n", v)
		noCapture := v
		go func() { results <- doWork(ctx, noCapture) }()
	}

	wr := make([]workResult, 0, numberOfItems)
	alreadyProcessed := 0
	for el := range results {
		fmt.Printf("alreadyProcessed: %v\n", alreadyProcessed)
		if el.err != nil {
			fmt.Printf("error: %v\n", el.err)
		} else {
			fmt.Printf("rcvd: %v\n", el.err)
			wr = append(wr, el)
		}

		alreadyProcessed += 1

		if alreadyProcessed == numberOfItems {
			close(results)
		}
	}
	fmt.Printf("wr: %v\n", wr)
}
