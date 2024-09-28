package main

import "fmt"

// Определяем структуру Person
type Person struct {
	Name     string
	lastName string
	Age      int
}

func main() {
	person := Person{Name: "Alice", lastName: "Alice", Age: 30} // Создаем экземпляр структуры
	person_2 := Person{Name: "Bob", lastName: "Bob", Age: 40}
	fmt.Println(person)                           // Вывод: {Alice 30}
	fmt.Println(person.Name, person.lastName)     // Вывод: Alice
	fmt.Println(person_2)                         // Вывод: {Alice 30}
	fmt.Println(person_2.Name, person_2.lastName) // Вывод: Alice
}
