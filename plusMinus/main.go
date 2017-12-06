package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//newLineReader returns a func that combines input.Scan() with input.Text() to make it one line
func newLineReader() func() string {
	input := bufio.NewScanner(os.Stdin)
	return func() string {
		input.Scan()
		return input.Text()
	}
}

//fMap is being used to create array of ints from strings
func fMap(og []string, f func(string) (int, error)) []int {
	vals := make([]int, len(og))
	for i, v := range og {
		vals[i], _ = f(v)
	}
	return vals
}

func main() {
	nextLine := newLineReader()
	t, _ := strconv.Atoi(nextLine())
	total := float32(t)
	nums := fMap(strings.Split(nextLine(), " "), strconv.Atoi)

	var p, n, z float32 = 0, 0, 0

	for _, val := range nums {
		if val > 0 {
			p++
		} else if val < 0 {
			n++
		} else {
			z++
		}
	}

	fmt.Printf("%f\n%f\n%f\n", p/total, n/total, z/total)
}
