package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "sepatator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Printf("\ndebug "+"n = %v", *n)
	fmt.Printf("\ndebug "+"sep = %v", *sep)
	fmt.Printf("\ndebug "+"flag.args = %v", flag.Args())

}
