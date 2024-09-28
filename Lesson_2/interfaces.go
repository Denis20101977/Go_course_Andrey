package main

import (
	"fmt"
	"math"
)

// Определяем интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Определяем тип Circle
type Circle struct {
	radius float64
}

// Реализуем метод Area для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Определяем тип Rectangle
type Rectangle struct {
	width, height float64
}

// Реализуем метод Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Функция для вычисления общей площади фигур
func totalArea(shapes []Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.Area() // Вызываем метод Area для каждой фигуры
	}
	return area
}

func main() {
	// Создаем массив фигур
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
	}

	// Вычисляем и выводим общую площадь
	fmt.Printf("Total Area: %.2fn", totalArea(shapes))
}
