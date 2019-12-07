package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type instruction struct {
	orig       int
	opcode     int
	m1, m2, m3 int
}

type program struct {
	memory  []int
	ip      int
	stdin   io.Reader
	stdout  io.Writer
	inChan  chan int
	outChan chan int
}

func NewProgram(memory []int, in chan int, out chan int) program {
	return program{memory: memory, ip: 0, inChan: in, outChan: out}
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
			//reader := bufio.NewReader(p.stdin)
			//s, _ := reader.ReadString('\n')
			//p.write(1, atoi(s[:len(s)-1]))
			p.write(1, <-p.inChan)
			p.ip += 2
		case 4:
			//writer := bufio.NewWriter(p.stdout)
			//writer.WriteString(fmt.Sprintf("%v", p.read(1, inst.m1)))
			//writer.Flush()
			p.outChan <- p.read(1, inst.m1)
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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

	inChan := make(chan int, 2)
	outChan := make(chan int, 1)
	var maxOutput int
	for _, phaseSettings := range permutations([]int{0, 1, 2, 3, 4}) {
		output := 0
		for _, setting := range phaseSettings {
			contentCopy := make([]int, len(content))
			copy(contentCopy, content)
			inChan <- setting
			inChan <- output
			p := NewProgram(contentCopy, inChan, outChan)
			p.Run()
			output = <-outChan
		}
		if output > maxOutput {
			maxOutput = output
		}
	}
	fmt.Println(maxOutput)

	maxOutput = 0
	for _, phaseSettings := range permutations([]int{5, 6, 7, 8, 9}) {
		c0 := make(chan int, 2)
		c1 := make(chan int, 2)
		c2 := make(chan int, 2)
		c3 := make(chan int, 2)
		c4 := make(chan int, 2)
		var contentCopy []int
		contentCopy = make([]int, len(content))
		copy(contentCopy, content)
		ampa := NewProgram(contentCopy, c0, c1)
		contentCopy = make([]int, len(content))
		copy(contentCopy, content)
		ampb := NewProgram(contentCopy, c1, c2)
		contentCopy = make([]int, len(content))
		copy(contentCopy, content)
		ampc := NewProgram(contentCopy, c2, c3)
		contentCopy = make([]int, len(content))
		copy(contentCopy, content)
		ampd := NewProgram(contentCopy, c3, c4)
		contentCopy = make([]int, len(content))
		copy(contentCopy, content)
		ampe := NewProgram(contentCopy, c4, c0)
		c0 <- phaseSettings[0]
		c1 <- phaseSettings[1]
		c2 <- phaseSettings[2]
		c3 <- phaseSettings[3]
		c4 <- phaseSettings[4]

		var wg sync.WaitGroup
		wg.Add(5)
		go func() {
			defer wg.Done()
			ampa.Run()
		}()
		go func() {
			defer wg.Done()
			ampb.Run()
		}()
		go func() {
			defer wg.Done()
			ampc.Run()
		}()
		go func() {
			defer wg.Done()
			ampd.Run()
		}()
		go func() {
			defer wg.Done()
			ampe.Run()
		}()
		c0 <- 0
		wg.Wait()
		output := <-c0
		if output > maxOutput {
			maxOutput = output
		}
	}
	fmt.Println(maxOutput)
}
