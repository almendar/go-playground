package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func process(val int) int {
	fmt.Printf("Processing value %d\n", val)
	val++
	return val
}

func runThingsConcurrently(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
		close(out)
	}()
}

func TestProcessing1(t *testing.T) {

	in := make(chan int)
	out := make(chan int)

	runThingsConcurrently(in, out)

	go func() {
		for i := 0; i < 10; i++ {
			c := i
			in <- c
		}
		close(in)
	}()
	for o := range out {
		fmt.Println(o)
	}
}

// func searchData(s string, searchers []func(string) []string) []string {
// 	done := make(chan struct{})
// 	result := make(chan []string)
// 	for _, searcher := range searchers {
// 		go func(searcher func(string) []string) {
// 			select {
// 			case result <- searcher(s):
// 			case <-done:
// 			}
// 		}(searcher)
// 	}
// 	r := <-result
// 	close(done)
// 	return r
// }

func TestDeadlock(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		fmt.Println("T1: Writting to ch1")
		ch1 <- v
		fmt.Println("T1: Written to ch1")
		fmt.Println("T1: Reading fomr ch2")
		v2 := <-ch2
		fmt.Println("T1: Read from ch2")
		fmt.Println(v, v2, "T1")
	}()

	v := 2

	var v2 int

loop:
	for {
		select {
		case v2 = <-ch1:
			fmt.Println("T2: Reading fomr ch2")
		case ch2 <- v:
			fmt.Println("T2: Written to ch2")
			break loop

		}
	}
	fmt.Println(v, v2, "T2")
	time.Sleep(1 * time.Second)
}

func TestCancelingFunction(t *testing.T) {

	type CancelFunc = func()

	countTo := func(max int) (<-chan int, CancelFunc) {
		ch := make(chan int)
		done := make(chan struct{})
		var cancel CancelFunc = func() {
			close(done)
		}
		go func() {
			for i := 0; i < max; i++ {
				select {
				case <-done:
					close(ch)
					return
				default:
					ch <- i
				}
			}
		}()
		return ch, cancel
	}

	readCh, cancel := countTo(100)
	var buffer [5]int
	for i := range readCh {
		if i >= 5 {
			cancel()
			break
		}
		buffer[i] = i
		fmt.Println(i, "!")
	}

	expect := [5]int{0, 1, 2, 3, 4}
	if buffer != expect {
		t.Errorf("expected: %v, got: %v", expect, buffer)
	}

}

func TestChannelWritting(t *testing.T) {
	ch := make(chan int, 1)
	i := 1
	ch <- i
	a := <-ch
	println(i, a)
}

func TestProcessingInParallel(t *testing.T) {

	processInParallel := func(parallelism int) []int {
		results := make(chan int, parallelism)
		for i := 0; i < parallelism; i++ {
			go func() {
				results <- process(i)
			}()
		}
		var out []int
		for i := 0; i < parallelism; i++ {
			out = append(out, <-results)
		}
		return out
	}

	out := processInParallel(5)

	for i, v := range out {
		if out[i] != v {
			t.Errorf("Values are wrong %v", out)
		}
	}
}

// -----------

// struct /
type PressureGauge struct {
	ch chan struct{}
}

func NewPressureGauge(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (g *PressureGauge) ProcessWithOk(callback func()) error {
	token, ok := <-g.ch
	if !ok {
		return fmt.Errorf("Too many requests")
	}
	fmt.Println("Not OK!")
	callback()
	g.ch <- token
	return nil
}

func (g *PressureGauge) ProcessWithSSP(callback func(id int), id int) error {
	select {
	case <-g.ch:
		callback(id)
		g.ch <- struct{}{}
		return nil
	default:
		fmt.Printf("DEFAULT %v\n", id)
		return fmt.Errorf("Too many requests: %v\n", id)
	}
}

func TestBackpressureGauge(t *testing.T) {

	sleep := func(id int) {
		fmt.Printf("Sleeping %d\n", id)
		time.Sleep(10 * time.Second)
	}

	parallelism := 50

	workStarted := make(chan int, parallelism)

	pressure := NewPressureGauge(parallelism)

	for i := 0; i < parallelism; i++ {
		id := i
		go func() {
			workStarted <- 1
			pressure.ProcessWithSSP(sleep, id)
		}()
	}

	startedRoutines := 0

	for i := range workStarted {
		startedRoutines += i
		if startedRoutines == parallelism {
			break
		}
	}

	err := pressure.ProcessWithSSP(sleep, 100)
	if err == nil {
		t.Error(err)
	}
}
