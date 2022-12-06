package main

import "fmt"

const (
	a     = 1
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

func ex4() {
	z := 234
	fmt.Printf("%d\t%b\t%#x\n", z, z, z)
	b := z << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

func ex5() {
	s := `raw string
lit
er
al
`
	fmt.Println(s)
}

func ex6() {
	const (
		d = 2022 + iota
		e = 2022 + iota
		f = 2022 + iota
		g = 2022 + iota
	)
	fmt.Println(d, e, f, g)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}
