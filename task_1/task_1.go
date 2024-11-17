package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, its my page: ")
	var message string
	fmt.Fscan(os.Stdin, &message)

	target := "http://"

	// Переменная для хранения результата
	var result []rune
	found := false

	for i, ch := range message {
		if !found && i <= len(message)-len(target) && message[i:i+len(target)] == target {
			found = true
			for _, t := range target {
				result = append(result, t) // Добавляем 'http://' в результат
			}
			// Заполняем остальную часть звёздочками
			for j := i + len(target); j < len(message); j++ {
				result = append(result, '*')
			}
			break // Прекращаем дальнейшую обработку, так как нашли наш target
		} else if !found {
			result = append(result, ch) // Добавляем символы до 'http://'
		}
	}

	if found {
		fmt.Println("Замаскированная строка:", string(result))
	} else {
		fmt.Println("Строка не содержит:", target)
	}

	fmt.Println("Hello, its my page: ", string(result), " See you")
}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Hello, its my page: ")
//	var message string
//	fmt.Scanln(&message)
//
//	// Создаем переменную для замаскированного сообщения
//	maskedMessage := ""
//
//	// Проходим по каждому символу в сообщении
//	for _, char := range message {
//		if char != ' ' { // Если символ не пробел
//			maskedMessage += "*" // Заменяем на '*'
//		} else {
//			maskedMessage += " " // Оставляем пробел без изменений
//		}
//	}
//
//	// Выводим замаскированное сообщение
//	fmt.Println("Hello, its my page: ", maskedMessage, " See you")
//}
