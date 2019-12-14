package main

import (
	"../utils"
	"../utils/intcode"
	"fmt"
	"time"
)

type XY struct {
	x, y int
}

type Arcade struct {
	canvas    map[XY]int
	computer  *intcode.Program
	display   int
	BallPos   XY
	PaddlePos XY
	done      chan bool
}

func NewArcade(code string) *Arcade {
	return &Arcade{canvas: make(map[XY]int), computer: intcode.ParseIntcodeProgram(code), done: make(chan bool)}
}

func (a *Arcade) Start(render bool) {
	go func() {
		defer func() {
			a.done <- true
		}()
		a.computer.Run()
	}()

	var done bool
	for !done {
		select {
		case x := <-a.computer.Out:
			y, t := <-a.computer.Out, <-a.computer.Out
			pos := XY{x, y}
			if pos == (XY{-1, 0}) {
				a.display = t
			} else {
				a.canvas[pos] = t
				if t == 4 {
					a.BallPos = pos
				} else if t == 3 {
					a.PaddlePos = pos
				}
			}
			if render && (t == 3 || t == 4) {
				a.Render()
			}
		case <-a.computer.InputNeeded:
			var input int
			if a.BallPos.x == a.PaddlePos.x {
				input = 0
			} else if a.BallPos.x > a.PaddlePos.x {
				input = 1
			} else {
				input = -1
			}
			if !a.computer.Done {
				a.computer.In <- input
			}
		case <-a.done:
			done = true
		}
	}
}

func findBounds(m map[XY]int) (XY, XY) {
	var minx, miny, maxx, maxy int
	for p := range m {
		if p.x < minx {
			minx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}
	return XY{minx, miny}, XY{maxx, maxy}
}

func (a *Arcade) Render() {
	fmt.Println("\033[2J")

	min, max := findBounds(a.canvas)
	for y := min.y; y <= max.y; y++ {
		fmt.Print("\n")
		for x := min.x; x <= max.x; x++ {
			switch a.canvas[XY{x, y}] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print(".")
			case 3:
				fmt.Print("-")
			case 4:
				fmt.Print("O")
			}
		}
	}
	fmt.Println("")
	fmt.Println(a.display)
	time.Sleep(2 * time.Millisecond)
}

func main() {
	lines := utils.ReadLines("input.txt")
	arcade := NewArcade(lines[0])
	render := false

	arcade.Start(render)

	var blockCount int
	for _, v := range arcade.canvas {
		if v == 2 {
			blockCount++
		}
	}
	fmt.Println(blockCount)

	arcade = NewArcade(lines[0])
	arcade.computer.Memory[0] = 2
	arcade.Start(render)
	fmt.Println(arcade.display)
}
