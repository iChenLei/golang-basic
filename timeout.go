package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Echo one message"
	}()

	select {
	case <-c1:
		fmt.Println("Ok ..")
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout 1")
	}
}
