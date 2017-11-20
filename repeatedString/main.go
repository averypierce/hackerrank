package main

//avery vankirk

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	input.Scan()
	n, _ := strconv.Atoi(input.Text())

	//first, count all the a's in s
	count := strings.Count(s, "a")
	//Then multiply for full repetitions of string
	total := count * int(math.Floor(float64(n/len(s))))
	//Add remainder
	total += strings.Count(s[:n%len(s)], "a")

	fmt.Println(total)
}
