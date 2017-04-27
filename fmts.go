package main

import (
	"fmt"
)

type point struct {
	x, y int
}

var f = fmt.Printf

func main() {
	p := point{1, 2}

	f("%v\n", p)
	f("%+v\n", p)
	f("%#v\n", p)
	f("%T\n", p)
	f("%b\n", 8)
	f("%p\n", &p)
}
