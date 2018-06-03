package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//$ curl https://github.com | ./findlinks1.exe
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	var cnt int
	for _, link := range myvisit(nil, doc, cnt) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func myvisit(links []string, n *html.Node, cnt int) []string {
	cnt++
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				links = append(links, fmt.Sprintf(" <----%d", cnt))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = myvisit(links, c, cnt)
	}

	return links
}
