package main

import (
	"fmt"
)

func main() {
	fmt.Println(faci(4))
}

func faci(n int) int {
	if n == 0 {
		return 1
	}
	return n * faci(n-1)
}
