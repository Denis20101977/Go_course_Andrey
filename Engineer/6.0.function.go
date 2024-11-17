package main

import (
  "fmt"
)

func addPrefix(origin string) string {
  return "prefix_" + origin
}

func newFunc() int {
  a := 10
  b := 10
  c := 10000
  return a + b*c // Пример использования переменных после их объявления
}

func addPrefixWithErr(origin string) (string, error) {
  return "prefix_with_err_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
  res = "prefix_" + origin
  length = len(res)
  return
}

// func main() {
  myString := addPrefix("my_string")
  fmt.Println(myString)

  myStringWithErr, err := addPrefixWithErr("error_string")
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println(myStringWithErr)

  myNew := newFunc()
  fmt.Println(myNew)
}
