package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//newLineReader returns a func that combines input.Scan() with input.Text() to make it one line
func newLineReader() func() string {
	input := bufio.NewScanner(os.Stdin)
	//Create "huge" buffer to make compatible with hackerranks larger inputs
	//Better: read word by word instead of line by line so we dont need the large buffer
	buf := make([]byte, 500000)
	input.Buffer(buf, 500000)
	return func() string {
		input.Scan()
		return input.Text()
	}
}

func ransomNote(mag, note []string) bool {

	m := len(mag)
	n := len(note)
	if m < n {
		return false
	}

	wordbank := make(map[string]int)

	for _, key := range mag {
		wordbank[key]++
	}
	for _, key := range note {
		wordbank[key]--
		if wordbank[key] < 0 {
			return false
		}
	}
	return true
}

func main() {
	nextLine := newLineReader()
	strings.Split(nextLine(), " ") //Skip first input line
	mag := strings.Split(nextLine(), " ")
	note := strings.Split(nextLine(), " ")

	if ransomNote(mag, note) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
