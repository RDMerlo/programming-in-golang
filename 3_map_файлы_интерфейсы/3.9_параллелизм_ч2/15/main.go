package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		array := make([]int, n)
		wg := new(sync.WaitGroup)

		for i := 0; i < n; i++ {
			wg.Add(1)
			x1 := <-in1
			x2 := <-in2

			go func(in1 <-chan int, in2 <-chan int, i int, wg *sync.WaitGroup) {
				defer wg.Done()

				r1 := make(chan int)
				r2 := make(chan int)

				defer close(r1)
				defer close(r2)

				go calc(fn, x1, r1)
				go calc(fn, x2, r2)

				array[i] = <-r1 + <-r2
			}(in1, in2, i, wg)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			out <- array[i]
		}
	}()
}

func calc(fn func(int) int, x int, res chan<- int) {
	res <- fn(x)
}
