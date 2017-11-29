package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ampm := input.Text()

	hours, _ := strconv.Atoi(ampm[:2])

	if hours == 12 {
		if ampm[8:9] == "A" {
			hours = 0
		}
	} else if hours < 12 && ampm[8:9] == "P" {
		hours += 12
	}

	fmt.Printf("%02d:%s", hours, ampm[3:8])

}
