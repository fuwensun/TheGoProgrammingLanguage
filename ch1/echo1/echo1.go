package main

import (
	"os";
	"fmt"
)

func main() {
	var s, sep string	//<----
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " -> "
	}
	fmt.Print(s)
}
