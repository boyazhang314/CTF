// Performs ROT13 on inputted text, prints the output

package main

import (
	"fmt"
	"strings"
)

func rot13(input string) string {
	cypher := make([]string, 0, len(input))

	for _, ch := range input {
		if 'a' <= ch && ch <= 'z' {
			ch = ((ch - 'a' + 13) % 26) + 'a'
		}
		if 'A' <= ch && ch <= 'Z' {
			ch = ((ch - 'A' + 13) % 26) + 'A'
		}
		cypher = append(cypher, string(ch))
	}

	output := strings.Join(cypher, "")
	return output
}

func main() {
	var input string
	fmt.Scanln(&input)

	fmt.Println(rot13(input))
}
