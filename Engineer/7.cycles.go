package main

import (
	"fmt"
)

// func main() {
	sl := []int64{9, 8, 7}
	for key, value := range sl {
		fmt.Printf("key: %v, val: %v \n", key, value)

	}

	for key, _ := range sl {
		fmt.Printf("value: %v \n", key)
	}

	ages := map[string]int{
		"Андрей":   30,
		"Светлана": 40,
	}
	for key, value := range ages {
		fmt.Printf("key: %v, val: %v \n", key, value)
	}

	word := "Денис"
	for i := 0; i < len(word); i++{
		fmt.Println(word[i])
		fmt.Printf("%T", word[i])
	}

}
