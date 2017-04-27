package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("Sync")

	go f("Async")

	go func(msg string) {
		fmt.Println(msg)
	}("Anonymous")

	fmt.Println("done")
}
