//package main
//
//import "fmt"
//
//func main() {
//	slice1 := []int{1, 2, 3}
//	slice1 = append(slice1, 4, 5, 6, 7)
//	fmt.Println(slice1)
//}

// При создании срезов с помощью make, вы можете указать длину и, по желанию, ёмкость среза.
package main

import "fmt"

func main() {
	// Создание среза целых чисел длиной 5 и ёмкостью 10
	slice := make([]int, 5, 10)

	// Инициализация значений
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1 // Заполнение значениями 1, 2, 3, 4, 5
	}

	fmt.Println("Срез:", slice)
	fmt.Println("Длина:", len(slice), "Ёмкость:", cap(slice))

	// Добавим еще элементы в срез
	slice = append(slice, 6, 7, 8)
	fmt.Println("После добавления:", slice)
	fmt.Println("Длина:", len(slice), "Ёмкость:", cap(slice))
}
