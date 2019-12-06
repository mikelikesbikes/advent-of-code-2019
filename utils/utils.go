package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("couldn't open %v", filename))
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
