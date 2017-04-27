package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	gitCmd := exec.Command("git")
	data, err := gitCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	time.Tick()
}
