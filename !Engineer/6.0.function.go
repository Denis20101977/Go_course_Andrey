package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_with_err_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return
}

func main() {
	myString := addPrefix("my_string")
	fmt.Println(myString)

	myString, err := addPrefixWithErr("error_string")
	fmt.Println(err)
	fmt.Println(myString)

}
