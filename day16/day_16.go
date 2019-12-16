package main

import (
	"../utils"
	"fmt"
	"strings"
)

var base = []int{0, 1, 0, -1}

func transform(digits []int, n int) int {
	var sum int
	for i, d := range digits {
		sum += d * base[((i+1)/n)%4]
	}
	return utils.Abs(sum) % 10
}

func FFT(input string, n int) string {
	if n == 0 {
		return input
	}

	digits := btoi([]byte(input))
	newDigits := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		newDigits[i] = transform(digits, i+1)
	}

	return FFT(string(itob(newDigits)), n-1)
}

func Decode(input string, n int) string {
	realSignal := strings.Repeat(input, 10000)
	offset := utils.Atoi(input[0:7])

	tailSig := btoi([]byte(realSignal[offset:]))

	for i := 0; i < n; i++ {
		var sum int
		for _, b := range tailSig {
			sum += b
		}

		var tailSigN []int
		for _, v := range tailSig {
			tailSigN = append(tailSigN, (utils.Abs(sum) % 10))
			sum -= v
		}
		tailSig = tailSigN
	}

	return string(itob(tailSig[:8]))
}

func btoi(arr []byte) []int {
	res := make([]int, len(arr))
	for i, b := range arr {
		res[i] = int(b) - 48
	}
	return res
}

func itob(arr []int) []byte {
	res := make([]byte, len(arr))
	for i, x := range arr {
		res[i] = byte(x + 48)
	}
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")

	fft := FFT(lines[0], 100)
	fmt.Println(fft[0:8])

	fmt.Println(Decode(lines[0], 100))
}
