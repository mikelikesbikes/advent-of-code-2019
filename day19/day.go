package main

import (
	"../utils"
	"../utils/intcode"
	"fmt"
)

type XY struct {
	x, y int
}

type Beam struct {
	code string
}

func NewBeam(code string) *Beam {
	return &Beam{code: code}
}

func (b *Beam) TestPos(pos XY) int {
	computer := intcode.ParseIntcodeProgram(b.code)
	go computer.Run()
	var toggle bool
	for {
		select {
		case <-computer.InputNeeded:
			if toggle {
				computer.In <- pos.y
			} else {
				computer.In <- pos.x
			}
			toggle = !toggle
		case c := <-computer.Out:
			return c
		}
	}
}

func (b *Beam) FitSquare(n int) XY {
	miny, maxy := 0, 100000
	var x int

	for miny < maxy {
		midy := miny + (maxy-miny)/2
		for x = 0; x < 100000; x++ {
			test := b.TestPos(XY{x, midy})
			if test == 1 {
				break
			}
		}
		tr := b.TestPos(XY{x + n - 1, midy - n + 1})
		if tr == 1 {
			maxy = midy
		} else {
			miny = midy + 1
		}
	}
	return XY{x, maxy - n + 1}
}

func main() {
	lines := utils.ReadLines("input.txt")

	beam := NewBeam(lines[0])

	var affected int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if beam.TestPos(XY{x, y}) == 1 {
				affected++
			}
		}
	}
	fmt.Println(affected)

	pos := beam.FitSquare(100)
	fmt.Println(pos.x*10000 + pos.y)
}
