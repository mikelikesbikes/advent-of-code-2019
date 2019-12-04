package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FuelCounterUpper(mass int) int {
	return (mass / 3) - 2
}

func FuelDoubleChecker(mass int) int {
	var totalFuel int
	fuel := FuelCounterUpper(mass)
	for fuel > 0 {
		totalFuel += fuel
		fuel = FuelCounterUpper(fuel)
	}
	return totalFuel
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open input.txt")
	}

	var masses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("couldn't convert text")
		}
		masses = append(masses, mass)
	}

	var p1 int
	for _, mass := range masses {
		p1 += FuelCounterUpper(mass)
	}

	fmt.Println(p1)

	var p2 int
	for _, mass := range masses {
		p2 += FuelDoubleChecker(mass)
	}

	fmt.Println(p2)
}
