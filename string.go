package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains  ", s.Contains("test", "es"))
	p("Count     ", s.Count("test", "t"))
	p("Hasprefix ", s.HasPrefix("testsss", "tes"))
	p("Index     ", s.Index("chenlei", "e"))
	p("Join      ", s.Join([]string{"a", "b"}, "_"))
	p("Split     ", s.Split("a-b-c-d", "-"))
	p("Touppercase", s.ToUpper("absbsshsj"))
}
