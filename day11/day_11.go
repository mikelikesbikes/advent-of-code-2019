package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	orig       int
	opcode     int
	m1, m2, m3 int
}

type Program struct {
	memory       []int
	ip           int
	relativeBase int
	in           chan int
	out          chan int
}

func NewProgram(memory []int) *Program {
	memoryCopy := make([]int, len(memory))
	copy(memoryCopy, memory)
	return &Program{memory: memoryCopy, ip: 0, in: make(chan int, 2), out: make(chan int, 20)}
}

func parseInst(i int) Instruction {
	pads := fmt.Sprintf("%05d", i)
	return Instruction{
		orig:   i,
		opcode: atoi(pads[3:]),
		m1:     atoi(string(pads[2])),
		m2:     atoi(string(pads[1])),
		m3:     atoi(string(pads[0])),
	}
}

func (p *Program) nextInst() Instruction {
	return parseInst(p.memory[p.ip])
}

func (p *Program) expandMemory(size int) {
	if len(p.memory)-1 < size {
		newMemory := make([]int, size+1)
		copy(newMemory, p.memory)
		p.memory = newMemory
	}
}

func (p *Program) paramVal(param, mode int) int {
	var val int
	pos := p.ip + param
	switch mode {
	case 0:
		// position mode
		position := p.memory[pos]
		p.expandMemory(position)
		val = p.memory[position]
	case 1:
		// immediate mode
		val = p.memory[pos]
	case 2:
		// relative
		relative := p.relativeBase + p.memory[pos]
		p.expandMemory(relative)
		val = p.memory[relative]
	default:
		panic(fmt.Sprintf("invalid parameter mode %v", mode))
	}
	return val
}

func (p *Program) write(param, mode, val int) {
	pos := p.ip + param
	p.expandMemory(pos)
	if mode == 2 {
		npos := p.relativeBase + p.memory[pos]
		pos = npos
	} else {
		pos = p.memory[pos]
	}
	p.expandMemory(pos)
	p.memory[pos] = val
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("can't convert %v", s))
	}
	return v
}

func (p *Program) Step() bool {
	inst := p.nextInst()
	switch inst.opcode {
	case 1:
		x, y := p.paramVal(1, inst.m1), p.paramVal(2, inst.m2)
		p.write(3, inst.m3, x+y)
		p.ip += 4
	case 2:
		x, y := p.paramVal(1, inst.m1), p.paramVal(2, inst.m2)
		p.write(3, inst.m3, x*y)
		p.ip += 4
	case 3:
		v := <-p.in
		p.write(1, inst.m1, v)
		p.ip += 2
	case 4:
		p.out <- p.paramVal(1, inst.m1)
		p.ip += 2
	case 5:
		x := p.paramVal(1, inst.m1)
		if x != 0 {
			p.ip = p.paramVal(2, inst.m2)
		} else {
			p.ip += 3
		}
	case 6:
		x := p.paramVal(1, inst.m1)
		if x == 0 {
			p.ip = p.paramVal(2, inst.m2)
		} else {
			p.ip += 3
		}
	case 7:
		x, y := p.paramVal(1, inst.m1), p.paramVal(2, inst.m2)
		lt := 0
		if x < y {
			lt = 1
		}
		p.write(3, inst.m3, lt)
		p.ip += 4
	case 8:
		x, y := p.paramVal(1, inst.m1), p.paramVal(2, inst.m2)
		eq := 0
		if x == y {
			eq = 1
		}
		p.write(3, inst.m3, eq)
		p.ip += 4
	case 9:
		val := p.paramVal(1, inst.m1)
		p.relativeBase += val
		p.ip += 2
	case 99:
		return true
	default:
		panic(fmt.Sprintf("unknown opcode %v", inst.orig))
	}
	return false
}

func (p *Program) Run() {
	for {
		halt := p.Step()
		if halt {
			close(p.in)
			close(p.out)
			return
		}
	}
}

type Coord struct {
	x, y int
}

type Painter struct {
	pos    Coord
	dir    int
	canvas map[Coord]int
	brain  *Program
}

func NewPainter(intcode []int) *Painter {
	return &Painter{
		pos:    Coord{0, 0},
		dir:    0,
		canvas: make(map[Coord]int),
		brain:  NewProgram(intcode),
	}
}

func (p *Painter) Start() {
	go func() {
		p.brain.Run()
	}()
}

func (p *Painter) Step() (res bool) {
	defer func() {
		if x := recover(); x != nil {
			res = false
		}
	}()
	p.brain.in <- p.canvas[p.pos]
	color := <-p.brain.out
	turn := <-p.brain.out
	p.canvas[p.pos] = color
	p.TurnAndMove(turn)
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
		p.pos.y += -1
	case 1:
		p.pos.x += 1
	case 2:
		p.pos.y += 1
	case 3:
		p.pos.x += -1
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
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	b := scanner.Text()
	var intcode []int
	for _, s := range strings.Split(string(b), ",") {
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
