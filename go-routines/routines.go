package main

import (
	"fmt"
	"time"
)

func count_nums() {
	for i := 0; i <= 10; i++ {
		time.Sleep(10 * time.Millisecond)
		println("num:", i)
	}
}

func main() {
	fmt.Println("Hello World!")
	go count_nums() //go keyword to get concurrency/parallelism : Instead of waiting for completion of this func, it executes the next func parallely.
	count_nums()
}
