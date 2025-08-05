package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	timer(3)
}

func timerSample() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done
	fmt.Println("all tasks are finished")
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func contextSample() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}

func timer(n int64) {
	fmt.Println("Timer has started.")
	ch := time.After(time.Duration(n) * time.Second)
	<-ch
	fmt.Println("Timer finished.")
}
