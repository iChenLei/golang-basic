package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job ", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for k := 1; k <= 3; k++ {
		jobs <- k
		fmt.Println("Send job ", k)
	}

	close(jobs)
	fmt.Println("Send all jobs...")
	kk:= <-done
	fmt.Println(kk)
}
