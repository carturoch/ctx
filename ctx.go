package main

import (
	"fmt"
	"os"
	s "strings"
)

// GetQuery reads the given arguments and
// concatenates them into a query string
func GetQuery(args []string) string {
	if len(args) < 2 {
		return ""
	}
	return s.Join(args[1:], " ")
}

func main() {
	val := GetQuery(os.Args)
	fmt.Println("Query: <", val, ">")
}
