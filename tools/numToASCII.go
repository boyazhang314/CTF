package main

import (
	"fmt"
)

func main() {
	str := ""

	var num int

	for true {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		str = str + string(num)
	}

	fmt.Println(str)
}
