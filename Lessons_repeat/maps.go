package main

import "fmt"

func main() {
	users := map[string]int{
		"Artem":  15,
		"Boris":  25,
		"Evgeny": 35,
	}
	users["Sergey"] = 21
	//delete(users, "Artem")

	for key, value := range users {
		fmt.Println(len(key), value)
	}
}
