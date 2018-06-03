package main

import (
	"TheGoProgrammingLanguage/ch6/intset"
	"fmt"
	/*
		."TheGoProgrammingLanguage/ch6/intset"
		_"TheGoProgrammingLanguage/ch6/intset"
		is"TheGoProgrammingLanguage/ch6/intset"
	*/
	. "TheGoProgrammingLanguage/ch6/geometry"
)

func main() {

	fmt.Printf("\n-----intset------")
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)

	fmt.Printf("\n-----geometry------")
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	perim1 := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	//fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
	fmt.Println(perim1.Distance()) // "12", method of geometry.Path

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	//p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	//Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal

	pptr.Distance(q)
	(*pptr).Distance(q)

	Point{1, 2}.Distance(q) //  Point
	pptr.ScaleBy(2)         // *Point

	p.ScaleBy(2) // implicit (&p)

	pptr.Distance(q) // implicit (*pptr)
}
