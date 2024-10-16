package main

import "fmt"

func main() {
	var list []int64

	l := len(list)
	c := cap(list)

	list = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	l = len(list)
	c = cap(list)

	//list = make([]int64, 0, 5)
	//l = len(list)
	//c = cap(list)

	list = append(list, 10)

	fmt.Println(list, l, c)

}
