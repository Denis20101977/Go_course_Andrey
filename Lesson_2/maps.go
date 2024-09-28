package main

import "fmt"

func main() {
	m := make(map[string]int) // Создаем карту, где ключи — строки, а значения — целые числа
	m["apple"] = 2            // Добавляем элемент
	m["banana"] = 3

	z := make(map[int]int)
	z[1] = 1
	z[2] = 2
	z[3] = 3
	z[4] = 4

	fmt.Println(m, z)                   // Вывод: map[apple:2 banana:3]
	fmt.Println(m["apple"], z[2], z[3]) // Вывод: 2 2 3
}
