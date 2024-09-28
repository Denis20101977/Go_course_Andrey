package main

import (
	"fmt"
	"math"
)

// Определим структуру Circle
type Circle struct {
	Radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Метод для вычисления периметра круга
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// Создаем экземпляр Circle
	myCircle := Circle{Radius: 5}

	// Используем методы
	fmt.Printf("Площадь круга: %.2fn", myCircle.Area())
	fmt.Printf("Периметр круга: %.2fn", myCircle.Perimeter())
}
