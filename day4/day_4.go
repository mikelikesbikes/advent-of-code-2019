package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ValidPassword(pw string) bool {
	return increasing(pw) && hasRepeat(pw)
}

func ValidPassword2(pw string) bool {
	return increasing(pw) && hasStandaloneDouble(pw)
}

func runner(f func(string) bool, min, max int) int {
	validChan := make(chan string)
	var wg sync.WaitGroup
	for i := min; i < max+1; i += 100 {
		mmax := i + 100
		if mmax > max {
			mmax = max + 1
		}
		wg.Add(1)
		go func(min, max int) {
			defer wg.Done()

			for ; min < max; min += 1 {
				pw := fmt.Sprintf("%v", min)
				if f(pw) {
					validChan <- pw
				}
			}
		}(i, mmax-1)
	}
	go func() {
		defer close(validChan)
		wg.Wait()
	}()

	count := 0
	for range validChan {
		count += 1
	}

	return count
}

func hasRepeat(pw string) bool {
	return pw[0] == pw[1] ||
		pw[1] == pw[2] ||
		pw[2] == pw[3] ||
		pw[3] == pw[4] ||
		pw[4] == pw[5]
}

func hasStandaloneDouble(pw string) bool {
	return (pw[0] == pw[1] && pw[1] != pw[2]) ||
		(pw[1] == pw[2] && pw[0] != pw[1] && pw[2] != pw[3]) ||
		(pw[2] == pw[3] && pw[1] != pw[2] && pw[3] != pw[4]) ||
		(pw[3] == pw[4] && pw[2] != pw[3] && pw[4] != pw[5]) ||
		(pw[4] == pw[5] && pw[3] != pw[4])
}

func increasing(pw string) bool {
	return pw[0] <= pw[1] &&
		pw[1] <= pw[2] &&
		pw[2] <= pw[3] &&
		pw[3] <= pw[4] &&
		pw[4] <= pw[5]
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %v to an int", s))
	}
	return i
}

func parseInput(s string) (int, int) {
	minmax := strings.Split(s, "-")
	if len(minmax) != 2 {
		panic(fmt.Sprintf("invalid input: %v", s))
	}
	return atoi(minmax[0]), atoi(minmax[1])
}

func main() {
	input := "138241-674034"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	min, max := parseInput(input)

	fmt.Println(runner(ValidPassword, min, max))
	fmt.Println(runner(ValidPassword2, min, max))
}
