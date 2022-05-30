package main

import "fmt"

func main() {
	var x int
	x = 42
	y := "James Bond"
	var z bool
	z = true
	fmt.Println(x, y, z)
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}