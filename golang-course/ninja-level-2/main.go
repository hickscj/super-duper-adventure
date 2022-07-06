package main

import "fmt"

const (
	a = 1
	b int = 2
)

func ex1() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}

func ex2() {
	g := 1 == 2
	h := 1 <= 2
	i := 1 >= 2
	j := 1 != 2
	k := 1 < 2
	l := 1 > 2
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

func ex3() {
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
}

func main() {
	ex1()
	ex2()
	ex3()
}