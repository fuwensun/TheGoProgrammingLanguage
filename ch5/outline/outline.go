package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)

	level := 0
	myoutline(nil, doc, level, 0)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func myoutline(stack []string, n *html.Node, depth int, count int) {
	depth++
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)

		front := strings.Repeat("-", depth)
		fmt.Printf("%s|depth=%3d,count=%3d|  +  %-6s  ", front, depth, count, n.Data)
		fmt.Println(append(stack, fmt.Sprintf(" <--%d", depth)))
	}

	var cnt int
	front := strings.Repeat("-", depth+1)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cnt == 0 {
			fmt.Printf("%s>depth=%3d  ###\n", front, depth+1)
		}
		myoutline(stack, c, depth, cnt)
		cnt++
	}
	if cnt > 0 {
		fmt.Printf("%s<depth=%3d  *** cnt=%3d  ***\n\n", front, depth+1, cnt)
	}
}
