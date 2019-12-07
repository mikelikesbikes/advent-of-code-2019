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
	memory []int
	ip     int
}

func NewProgram(memory []int) program {
	return program{memory: memory, ip: 0}
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

func (p program) nextInst() instruction {
	return parseInst(p.memory[p.ip])
}

func (p program) read(param, mode int) int {
	var val int
	pos := p.ip + param
	if mode == 0 {
		val = p.memory[p.memory[pos]]
	} else {
		val = p.memory[pos]
	}
	return val
}

func (p program) write(param, val int) {
	pos := p.ip + param
	p.memory[p.memory[pos]] = val
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("can't convert %v", s))
	}
	return v
}

func (p program) Run() {
	for p.ip < len(p.memory) {
		inst := p.nextInst()
		switch inst.opcode {
		case 1:
			x, y := p.read(1, inst.m1), p.read(2, inst.m2)
			p.write(3, x+y)
			p.ip += 4
		case 2:
			x, y := p.read(1, inst.m1), p.read(2, inst.m2)
			p.write(3, x*y)
			p.ip += 4
		case 3:
			reader := bufio.NewReader(os.Stdin)
			s, _ := reader.ReadString('\n')
			p.write(1, atoi(s[:len(s)-1]))
			p.ip += 2
		case 4:
			fmt.Println(p.read(1, inst.m1))
			p.ip += 2
		case 5:
			x := p.read(1, inst.m1)
			if x != 0 {
				p.ip = p.read(2, inst.m2)
			} else {
				p.ip += 3
			}
		case 6:
			x := p.read(1, inst.m1)
			if x == 0 {
				p.ip = p.read(2, inst.m2)
			} else {
				p.ip += 3
			}
		case 7:
			x, y := p.read(1, inst.m1), p.read(2, inst.m2)
			lt := 0
			if x < y {
				lt = 1
			}
			p.write(3, lt)
			p.ip += 4
		case 8:
			x, y := p.read(1, inst.m1), p.read(2, inst.m2)
			eq := 0
			if x == y {
				eq = 1
			}
			p.write(3, eq)
			p.ip += 4
		case 99:
			return
		default:
			panic(fmt.Sprintf("unknown opcode %v", inst.orig))
		}
	}
	return
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
	p := NewProgram(content)
	p.Run()
}
