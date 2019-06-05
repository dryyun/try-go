package main

import "fmt"

type Point struct {
	X, Y int
	Z    string
}

func main() {
	type Point struct {
		X, Y int
		Z    string
	}
	var p1 Point
	fmt.Printf("%T\t%#v\n", p1, p1) // main.Point      main.Point{X:0, Y:0, Z:""}

	p2 := Point{}
	fmt.Printf("%T\t%#v\n", p2, p2) // main.Point      main.Point{X:0, Y:0, Z:""}

	p3 := new(Point)
	fmt.Printf("%T\t%p\t%#v\n", p3, p3, p3) // *main.Point     0xc00000c080    &main.Point{X:0, Y:0, Z:""}
	fmt.Printf("%T\t%#v\n", *p3, *p3)       // main.Point      main.Point{X:0, Y:0, Z:""}
}
