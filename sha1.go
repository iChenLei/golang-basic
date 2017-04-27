package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "my name is chenlei"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
