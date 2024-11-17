package main

import "fmt"

func main() {
	var x1 [5]int
	var x2 [0]int
	var x3 [0]string

	var arr = [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}
	arr5 := [...]int{1, 2, 3}
	s := len(arr5)

	fmt.Println(x1, x2, x3, arr, arr3, arr5, arr3, s)
	fmt.Println(s)
}
