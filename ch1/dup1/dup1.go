package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		//fmt.Println(">>>  "+input.Text())
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("-->  %d\t%s\n", n, line)
			//fmt.Println("-->  %d\t%s\n", n, line)
			//fmt.Print("--> ", n, "===", line,"\n")
		}
	}

	fmt.Printf("<--the end--> %d\n", len(counts))
}
