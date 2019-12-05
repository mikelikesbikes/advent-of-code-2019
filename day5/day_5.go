package main

import (
	"fmt"
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
}
