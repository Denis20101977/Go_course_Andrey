В языке программирования Go (Golang) есть несколько основных типов данных. Вот краткий обзор каждого из них с примерами:

### 1. Булевый тип (bool)
Используется для хранения значений true или false.

package main

import "fmt"

func main() {
var isActive bool = true
fmt.Println("Is Active:", isActive)
}


### 2. Целочисленные типы (int, int8, int16, int32, int64)
Типы для хранения целых чисел различной разрядности.

package main

import "fmt"

func main() {
var a int = 42
var b int8 = -128
var c int16 = 30000
var d int32 = 2000000000
var e int64 = 9223372036854775807

fmt.Println(a, b, c, d, e)
}


### 3. Беззнаковые целочисленные типы (uint, uint8, uint16, uint32, uint64)
Используются для хранения положительных целых чисел.

package main

import "fmt"

func main() {
var a uint = 42
var b uint8 = 255

fmt.Println(a, b)
}


### 4. Числа с плавающей запятой (float32, float64)
Типы для хранения чисел с плавающей запятой.

package main

import "fmt"

func main() {
var x float32 = 3.14
var y float64 = 2.718281828459045

fmt.Println(x, y)
}


### 5. Комплексные числа (complex64, complex128)
Используются для хранения комплексных чисел.

package main

import "fmt"

func main() {
var c1 complex64 = 1 + 2i
var c2 complex128 = 3 + 4i

fmt.Println(c1, c2)
}


### 6. Символы (rune)
Тип для хранения символов (один символ в кодировке UTF-8).

package main

import "fmt"

func main() {
var ch rune = 'G'
fmt.Println("Rune:", ch)
}


### 7. Строки (string)
Для хранения строковых данных.

package main

import "fmt"

func main() {
var greeting string = "Hello, World!"
fmt.Println(greeting)
}


### 8. Массивы
Фиксированные массивы с элементами одного типа.

package main

import "fmt"

func main() {
var arr [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(arr)
}


### 9. Срезы (slise)
Динамические массивы.

package main

import "fmt"

func main() {
slice := []int{1, 2, 3, 4, 5}
fmt.Println(slice)
}


### 10. Карты (map)
Ассоциативные массивы для хранения пар "ключ-значение".

package main

import "fmt"

func main() {
ages := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(ages)
}


### 11. Структуры (struct)
Пользовательские типы данных, объединяющие разные типы.

package main

import "fmt"

type Person struct {
Name string
Age  int
}

func main() {
p := Person{Name: "Alice", Age: 25}
fmt.Println(p)
}


### 12. Интерфейсы (interface)
Тип данных, который может представлять любой тип.

package main

import "fmt"

type Animal interface {
Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
return "Woof!"
}

func main() {
var a Animal = Dog{}
fmt.Println(a.Speak())
}


Эти примеры демонстрируют основные типы данных в Go, а также некоторые связанные с ними конструкции, 
которые часто используются в разработке.
