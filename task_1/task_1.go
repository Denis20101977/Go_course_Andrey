// package main
//
// import (
//
//	"fmt"
//
// )
//
// func main() {
//
//		fmt.Println("Hello, its my page: ")
//		var message string
//		fmt.Scanln(&message)
//
//		// Создаем переменную для поиска в слайсе
//		//http := "http://"
//
//		// Преобразуем строку в байтовый слайс
//		bytes := []byte(message)
//		fmt.Println("Байты: ", bytes)
//
//		// Создаем цикл для поиска https://www.
//
//		// Преобразуем байты обратно в строку и выводим
//		convertedMessage := string(bytes)
//		fmt.Println("Hello, its my page: ", convertedMessage, "See you")
//	}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, its my page: ")
	var message string
	fmt.Scanln(&message)

	// Создаем переменную для замаскированного сообщения
	maskedMessage := ""

	// Проходим по каждому символу в сообщении
	for _, char := range message {
		if char != ' ' { // Если символ не пробел
			maskedMessage += "*" // Заменяем на '*'
		} else {
			maskedMessage += " " // Оставляем пробел без изменений
		}
	}

	// Выводим замаскированное сообщение
	fmt.Println("Hello, its my page: ", maskedMessage, " See you")
}
