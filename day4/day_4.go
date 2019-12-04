package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidPassword(pw string) bool {
	if increasing(pw) && hasRepeat(pw) {
		return true
	}
	return false
}

func ValidPassword2(pw string) bool {
	if increasing(pw) && hasStandaloneDouble(pw) {
		return true
	}
	return false
}

func hasRepeat(pw string) bool {
	if pw[0] == pw[1] ||
		pw[1] == pw[2] ||
		pw[2] == pw[3] ||
		pw[3] == pw[4] ||
		pw[4] == pw[5] {
		return true
	}
	return false
}

func hasStandaloneDouble(pw string) bool {
	if (pw[0] == pw[1] && pw[1] != pw[2]) ||
		(pw[1] == pw[2] && pw[0] != pw[1] && pw[2] != pw[3]) ||
		(pw[2] == pw[3] && pw[1] != pw[2] && pw[3] != pw[4]) ||
		(pw[3] == pw[4] && pw[2] != pw[3] && pw[4] != pw[5]) ||
		(pw[4] == pw[5] && pw[3] != pw[4]) {
		return true
	}
	return false
}

func increasing(pw string) bool {
	if pw[0] <= pw[1] &&
		pw[1] <= pw[2] &&
		pw[2] <= pw[3] &&
		pw[3] <= pw[4] &&
		pw[4] <= pw[5] {
		return true
	}
	return false
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %v to an int", s))
	}
	return i
}

func main() {
	input := "138241-674034"
	if arg1 := os.Args[1]; arg1 != "" {
		input = arg1
	}
	minmax := strings.Split(input, "-")
	if len(minmax) != 2 {
		panic(fmt.Sprintf("invalid input: %v", input))
	}
	min, max := atoi(minmax[0]), atoi(minmax[1])
	count := 0
	for i := min; i <= max; i++ {
		pw := fmt.Sprintf("%v", i)
		if ValidPassword(pw) {
			count += 1
		}
	}
	fmt.Println(count)

	count = 0
	for i := min; i <= max; i++ {
		pw := fmt.Sprintf("%v", i)
		if ValidPassword2(pw) {
			count += 1
		}
	}
	fmt.Println(count)
}
