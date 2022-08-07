package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var str string

	for true {
		_, err := fmt.Scanln(&str)
		if err != nil {
			break
		}

		text, err := hex.DecodeString(str)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(text))
	}
}
