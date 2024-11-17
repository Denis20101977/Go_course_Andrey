package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}

	p := Point{
		X: 5,
		Y: 15,
	}

	p.X = 68

	fmt.Println(p.X)
	fmt.Println(p.Y)

}