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

func (inst instruction) arity() int {
	var arity int
	switch inst.opcode {
	case 3, 4:
		arity = 1
	case 5, 6:
		arity = 2
	case 1, 2, 7, 8:
		arity = 3
	case 99:
		arity = 99
	}
	return arity
}

func (p program) read(param, mode int) int {
	var val int
	if mode == 0 {
		val = p.memory[p.memory[param]]
	} else {
		val = p.memory[param]
	}
	return val
}

func (p program) write(param, val int) {
	p.memory[p.memory[param]] = val
}

func (p1 program) dup() program {
	var p2 program
	p2.memory = make([]int, len(p1.memory))
	copy(p2.memory, p1.memory)
	return p2
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("can't convert %v", s))
	}
	return v
}

func (p program) Run() program {
	result := p.dup()
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(result.memory); {
		inst := parseInst(result.read(i, 1))
		jump := false
		switch inst.opcode {
		case 1:
			x, y := result.read(i+1, inst.m1), result.read(i+2, inst.m2)
			result.write(i+3, x+y)
		case 2:
			x, y := result.read(i+1, inst.m1), result.read(i+2, inst.m2)
			result.write(i+3, x*y)
		case 3:
			// read from stdin
			s, _ := reader.ReadString('\n')
			result.write(i+1, atoi(s[:len(s)-1]))
		case 4:
			fmt.Println(result.read(i+1, inst.m1))
		case 5:
			x := result.read(i+1, inst.m1)
			if x != 0 {
				jump = true
				i = result.read(i+2, inst.m2)
			}
		case 6:
			x := result.read(i+1, inst.m1)
			if x == 0 {
				jump = true
				i = result.read(i+2, inst.m2)
			}
		case 7:
			x, y := result.read(i+1, inst.m1), result.read(i+2, inst.m2)
			lt := 0
			if x < y {
				lt = 1
			}
			result.write(i+3, lt)
		case 8:
			x, y := result.read(i+1, inst.m1), result.read(i+2, inst.m2)
			eq := 0
			if x == y {
				eq = 1
			}
			result.write(i+3, eq)
		case 99:
			return result
		default:
			panic(fmt.Sprintf("unknown opcode %v", inst.orig))
		}
		if !jump {
			i += (inst.arity() + 1)
		}
	}
	return result
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
	p := program{memory: content}
	result := p.Run()
	fmt.Println(result)
}
