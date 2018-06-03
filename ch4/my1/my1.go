package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func AddStr(list string) { m[list]++ }

func main() {

	a := [...]string{"a", "b", "c"}
	a1 := k(a[:])
	a2 := k(a[:1])
	a3 := k(a[1:])

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Printf("-> %T\n", a1)
	fmt.Printf("-> %T\n", a2)
	fmt.Printf("-> %T\n", a3)

	str0 := "[\"a\" \"b\" \"c\"]"
	str1 := `["a" "b" "c"]`
	str2 := "abc"
	fmt.Printf("--> %T\n", str1)
	fmt.Printf("--> %v\n", str1)
	fmt.Printf("--> %q\n", str1)

	fmt.Printf("---> %T\n", str2)
	fmt.Printf("---> %v\n", str2)
	fmt.Printf("---> %q\n", str2)

	fmt.Printf("****>%t\n", a1 == str0)
	fmt.Printf("****>%t\n", a1 == str1)

	s1 := a[:]
	s2 := a[:1]
	s3 := a[1:]

	Add(s1)
	Add(s2)
	Add(s3)
	AddStr(str0)
	AddStr(str1)

	fmt.Println(Count(s1))
	fmt.Println(Count(s2))
	fmt.Println(Count(s3))
}
