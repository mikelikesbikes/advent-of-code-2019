package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	orig       int
	opcode     int
	m1, m2, m3 int
}

type Program struct {
	Memory       []int
	ip           int
	relativeBase int
	InputNeeded  chan bool
	In           chan int
	Out          chan int
	Done         bool
}

func ParseIntcodeProgram(code string) *Program {
	var intcode []int
	for _, s := range strings.Split(string(code), ",") {
		c, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("couldn't convert text: %v", s))
		}
		intcode = append(intcode, c)
	}
	return NewProgram(intcode)
}

func NewProgram(memory []int) *Program {
	memoryCopy := make([]int, len(memory))
	copy(memoryCopy, memory)
	return &Program{Memory: memoryCopy, ip: 0, In: make(chan int), Out: make(chan int), InputNeeded: make(chan bool), Done: false}
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
	return parseInst(p.Memory[p.ip])
}

func (p *Program) expandMemory(size int) {
	if len(p.Memory)-1 < size {
		newMemory := make([]int, size+1)
		copy(newMemory, p.Memory)
		p.Memory = newMemory
	}
}

func (p *Program) paramVal(param, mode int) int {
	var val int
	pos := p.ip + param
	switch mode {
	case 0:
		// position mode
		position := p.Memory[pos]
		p.expandMemory(position)
		val = p.Memory[position]
	case 1:
		// immediate mode
		val = p.Memory[pos]
	case 2:
		// relative
		relative := p.relativeBase + p.Memory[pos]
		p.expandMemory(relative)
		val = p.Memory[relative]
	default:
		panic(fmt.Sprintf("invalid parameter mode %v", mode))
	}
	return val
}

func (p *Program) write(param, mode, val int) {
	pos := p.ip + param
	p.expandMemory(pos)
	if mode == 2 {
		npos := p.relativeBase + p.Memory[pos]
		pos = npos
	} else {
		pos = p.Memory[pos]
	}
	p.expandMemory(pos)
	p.Memory[pos] = val
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
		p.InputNeeded <- true
		v := <-p.In
		p.write(1, inst.m1, v)
		p.ip += 2
	case 4:
		p.Out <- p.paramVal(1, inst.m1)
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
		p.Done = true
		return true
	default:
		panic(fmt.Sprintf("unknown opcode %v", inst.orig))
	}
	return false
}

func (p *Program) Run() {
	//defer close(p.In)
	defer close(p.Out)
	defer close(p.InputNeeded)

	for !p.Done {
		p.Step()
	}
}
