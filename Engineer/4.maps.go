package main

import "fmt"

// func main() {

	var m1 map[int64]bool
	var m2 map[string]string
	m3 := make(map[string]int)

	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 50,
	}

	age := ages["Андрей"]
	ages["Наталья"] = 31

	
	fmt.Println(ages)
	fmt.Println(ages["Иван"], 25)
	fmt.Println(m1, m2, m3, age)

}
