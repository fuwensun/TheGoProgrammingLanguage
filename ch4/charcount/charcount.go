package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for{
		r, n, err := in.ReadRune()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1{
			invalid++
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts{
		fmt.Printf("%q\t%d\n",c,n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen{
		fmt.Printf("%d\t%d\n",i,n)
	}

	if invalid > 0{
		fmt.Printf("\n%d invalid UTF-8 characters\n",invalid)
	}

//	加入汉字，为了测试
}
