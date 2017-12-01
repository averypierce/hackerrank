package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	input := bufio.NewReader(os.Stdin)
	n := 1

	for {
		c, _, err := input.ReadRune()
		if err == nil {
			if unicode.IsUpper(c) {
				n++
			}
		} else {
			break
		}

	}
	fmt.Println(n)

}
