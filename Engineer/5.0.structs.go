package main

import "fmt"

// func main() {
	type Point struct {
		X int
		Y int
		Z string
	}

	p := Point{
		X: 5,
		Y: 15,
		Z: "27",
	}

	p.X = 68
	p.Y = 46
	p.Z = "Hello"
	p.Z = "Hello"

	fmt.Println(p.X, p.Y)
	fmt.Println(p.Y, p.Z)

}
