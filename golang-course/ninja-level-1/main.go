/*
	EXERCISES
*/
package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

type spaceship string
type hotdog int

var ncc spaceship = "Enterprise"
var d hotdog = 5

const (
	a = iota
	b 
	c 
)

func main() {
	var s string = fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
	fmt.Printf("%v\t%T\n", ncc, ncc)
	fmt.Printf("%v\t%T\n", d, d)
	l := int(d)
	fmt.Println(l)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%#U", 3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}