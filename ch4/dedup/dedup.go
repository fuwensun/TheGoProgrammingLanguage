package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2 条测试指令
//	du -h -d2 / > mytestfile
//	du -h -d2 / >> mytestfile
//	./dedup.exe < ./mytestfile
//	du -h -d2 / | ./dedup.exe

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println("-->" + line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
