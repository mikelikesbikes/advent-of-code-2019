package main

import (
	"../utils"
	"../utils/intcode"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Painter struct {
	pos    Coord
	dir    int
	canvas map[Coord]int
	brain  *intcode.Program
}

func NewPainter(code []int) *Painter {
	return &Painter{
		pos:    Coord{0, 0},
		dir:    0,
		canvas: make(map[Coord]int),
		brain:  intcode.NewProgram(code),
	}
}

func (p *Painter) Start() {
	go p.brain.Run()
}

func (p *Painter) Step() (res bool) {
	defer func() {
		// attempting to read from a closed channel signals that the program is
		// finished and step should return false to tell the painter to halt
		if x := recover(); x != nil {
			res = false
		}
	}()

	p.brain.In <- p.canvas[p.pos]
	p.canvas[p.pos] = <-p.brain.Out
	p.TurnAndMove(<-p.brain.Out)
	return true
}

func (p *Painter) TurnAndMove(turn int) {
	switch turn {
	case 0:
		p.dir += 3
	case 1:
		p.dir += 1
	}

	switch p.dir & 3 {
	case 0:
		p.pos.y--
	case 1:
		p.pos.x++
	case 2:
		p.pos.y++
	case 3:
		p.pos.x--
	}
}

func Render(canvas map[Coord]int) {
	// find bounds
	var minx, miny, maxx, maxy int
	for k := range canvas {
		x, y := k.x, k.y
		if x < minx {
			minx = x
		}
		if y < miny {
			miny = y
		}
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	white := color.New(color.BgWhite).PrintfFunc()
	black := color.New(color.BgBlack).PrintfFunc()
	for y := miny; y <= maxy; y++ {
		fmt.Println()
		for x := minx; x <= maxx; x++ {
			switch canvas[Coord{x, y}] {
			case 0:
				black(" ")
			case 1:
				white(" ")
			}
		}
	}
}

func main() {
	line := utils.ReadLines("input.txt")[0]
	var intcode []int
	for _, s := range strings.Split(string(line), ",") {
		c, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("couldn't convert text: %v", s))
		}
		intcode = append(intcode, c)
	}
	painter := NewPainter(intcode)
	painter.Start()
	for painter.Step() {
	}

	fmt.Println(len(painter.canvas))

	painter = NewPainter(intcode)
	painter.canvas[painter.pos] = 1
	painter.Start()
	for painter.Step() {
	}
	Render(painter.canvas)
}
