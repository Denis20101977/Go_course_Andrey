package main

import "fmt"

func main() {
	// var list []int64

	

	list := []int64{1, 2, 3, 4, 5}
	l :=len(list)
	c :=len(list)

	list2 := make([]int64, 25, 50)
	list3 := make([]string, 25, 50)

	list2 = append(list2, 22)


	fmt.Println(l, c, list2, len(list3), list3)
	
}
