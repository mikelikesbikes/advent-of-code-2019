package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Compute(computer []int) []int {
	result := make([]int, len(computer))
	copy(result, computer)

	for i := 0; i < len(computer); i += 4 {
		switch result[i] {
		case 1:
			result[result[i+3]] = result[result[i+1]] + result[result[i+2]]
		case 2:
			result[result[i+3]] = result[result[i+1]] * result[result[i+2]]
		case 99:
			return result
		default:
			panic(fmt.Sprintf("unknown opcode %v", result[i]))
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open input.txt")
	}

	b, err := ioutil.ReadAll(file)
	var content []int
	for _, s := range strings.Split(string(b), ",") {
		c, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("couldn't convert text: %v", s))
		}
		content = append(content, c)
	}
	content[1] = 12
	content[2] = 2
	result := Compute(content)
	fmt.Println(result[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			content[1] = noun
			content[2] = verb
			result := Compute(content)
			if result[0] == 19690720 {
				fmt.Println((100 * noun) + verb)
				return
			}
		}
	}
}
