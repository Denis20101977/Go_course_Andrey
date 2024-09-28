package main

import (
	"fmt"
	"strconv"
)

func main() {
	group := map[string]string{"1": "1", "2": "2"}
	fmt.Println(group[strconv.Itoa(1)])

}
