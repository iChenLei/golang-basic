package main

import (
	"encoding/json"
	"fmt"
	"os"
	"encoding/xml"
)

type rs struct {
	Page   int
	Fruits []string
}

// type rss struct{
// 	Page int `json:"page"`
// 	Fruits []string `json:"fruits"`
// }

func main() {
	rs1 := &rs{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	rs1s, _ := json.Marshal(rs1)
	fmt.Println(string(rs1s))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)	
}
