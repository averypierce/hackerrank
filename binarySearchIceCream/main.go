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

//ts is trip to store
func ts(money int, a []int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		complement := money - a[i]

		//check if m[] contains the key 'complement'
		_, ok := m[complement]
		if ok {
			return m[complement] + 1, i + 1
		}
		m[a[i]] = i
	}
	//This line should never be reached, Error state, i guess
	return -1, -1
}

func main() {
	nextLine := newLineReader()
	t, _ := strconv.Atoi(nextLine())

	for i := 0; i < t; i++ {
		m, _ := strconv.Atoi(nextLine())
		//n, _ := strconv.Atoi(nextLine())
		nextLine()

		a := fMap(strings.Split(nextLine(), " "), strconv.Atoi)
		f, l := ts(m, a)
		fmt.Println(f, l)
	}
}
