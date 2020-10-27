package main

import (
	"fmt"
	"time"
)

//计算斐波那契数列并写到ch中
func fibonacci(n int, ch chan<- int) {
	first, second := 0, 1
	for i := 0; i < n; i++ {
		ch <- first
		first, second = second, first+second
	}
	close(ch)
}

func main() {
	ch := make(chan int, 40)
	i := 0
	start := time.Now()
	go fibonacci(cap(ch), ch)
	for result := range ch {
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		i++
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("took the time: %s\n", delta)
}
