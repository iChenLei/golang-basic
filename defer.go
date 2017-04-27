package main

import (
	"fmt"
	"os"
)

func main() {
	f := createfile("E:\\py\\defer.txt")
	defer closefile(f)
	writefile(f)
}

func createfile(p string) *os.File {
	fmt.Println("Creating a file...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writefile(f *os.File) {
	fmt.Println("Append data to file")
	fmt.Fprintln(f, "A file created by golang")
}

func closefile(f *os.File) {
	fmt.Println("Close file io writer")
	f.Close()
}
