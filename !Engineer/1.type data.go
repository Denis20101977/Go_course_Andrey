package main

import (
	"fmt"
)

func main() {
	//var myString string
	//var hello string = "\tHello\n"
	//
	//word := "my string"
	//
	//var b byte = 'b'
	//var r rune = 'r'

	str := "123456789"
	fmt.Println(str[5])
	fmt.Println(str[3])
	fmt.Println(str[8])
	_ = len(str)
	fmt.Println(len(str))

	newStr := str[6:]
	newStr = str[7:]
	fmt.Println(newStr)

	words := "Hello"
	hello := words + " World"
	fmt.Println(hello)

	//fmt.Println(myString, hello, word, r, b, word)

}
