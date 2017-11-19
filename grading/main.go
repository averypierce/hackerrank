package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Using custom type grade just for fun
type grade int

func main() {

	inputs := bufio.NewScanner(os.Stdin)
	//skip first line
	inputs.Scan()
	for inputs.Scan() {
		i, _ := strconv.Atoi(inputs.Text())
		finalGrade := calculate(grade(i))
		fmt.Println(finalGrade)
	}
}

func calculate(g grade) grade {
	if g < 0 || g > 100 {
		//error?
	}
	if g >= 38 {
		g = g.roundUp()
	}
	return g
}

//roundUp is a grade method
func (g grade) roundUp() grade {
	diff := 5 - g%5
	if diff < 3 {
		g += diff
	}
	return g
}
