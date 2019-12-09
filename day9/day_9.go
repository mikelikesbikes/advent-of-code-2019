package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	orig       int
	opcode     int
	m1, m2, m3 int
}

type program struct {
	memory       []int
	ip           int
	relativeBase int
	in           chan int
	out          chan int
}

func NewProgram(memory []int) *program {
	memoryCopy := make([]int, len(memory))
	copy(memoryCopy, memory)
	return &program{memory: memoryCopy, ip: 0, in: make(chan int, 2), out: make(chan int, 20)}
}

func parseInst(i int) instruction {
	pads := fmt.Sprintf("%05d", i)
	return instruction{
		orig:   i,
		opcode: atoi(pads[3:]),
		m1:     atoi(string(pads[2])),
		m2:     atoi(string(pads[1])),
		m3:     atoi(string(pads[0])),
	}
}

func (p *program) nextInst() instruction {
	return parseInst(p.memory[p.ip])
}

func (p *program) expandMemory(size int) {
	if len(p.memory)-1 < size {
		newMemory := make([]int, size+1)
		copy(newMemory, p.memory)
		p.memory = newMemory
	}
}

func (p *program) paramVal(param, mode int) int {
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

func (p *program) write(param, mode, val int) {
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

func (p *program) Step() bool {
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

func (p *program) Run() {
	for {
		halt := p.Step()
		if halt {
			return
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
	var content []int
	for _, s := range strings.Split(string(b), ",") {
		c, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("couldn't convert text: %v", s))
		}
		content = append(content, c)
	}
	p1 := NewProgram(content)
	p1.in <- 1
	go func() {
		defer close(p1.out)
		p1.Run()
	}()
	for i := range p1.out {
		fmt.Println(i)
	}

	p2 := NewProgram(content)
	p2.in <- 2
	go func() {
		defer close(p2.out)
		p2.Run()
	}()
	for i := range p2.out {
		fmt.Println(i)
	}
}
