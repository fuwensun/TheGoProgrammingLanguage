package main

import "fmt"

var m = make(map[int]int)

func main() {

	m[0] = 0
	fmt.Printf("--> %d map %v\n", 0, m[0])

	fmt.Printf("--> %d map %v\n", 1, m[1])

	v, ok := m[3]
	fmt.Printf("--> v = %v  ok = %v\n", v, ok)

	m[0] = 10
	m[1] = 11
	m[2] = 12
	m[5] = 15

	for i := 0; i < 10; i++ {
		v, ok := m[i] //v,ok
		fmt.Printf("--2> v = %v  ok = %v\n", v, ok)
	}
	for i := 0; i < 10; i++ {
		v = m[i] //v
		fmt.Printf("--3> v = %v\n", v)
	}

	for k, v := range m { //k,v
		fmt.Printf("--4> k = %v  v = %v\n", k, v)
	}
	for k := range m { //k
		fmt.Printf("--5> k = %v\n", k)
	}
}
