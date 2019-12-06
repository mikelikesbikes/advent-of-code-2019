package main

import (
	"../utils"
	"fmt"
	"strings"
)

type UOM struct {
	m map[string]string
}

func (uom *UOM) pathToCOM(s string) []string {
	var path []string
	for s != "COM" {
		s = uom.m[s]
		path = append(path, s)
	}
	return path
}

func (uom *UOM) checksumAt(s string) int {
	return len(uom.pathToCOM(s))
}

func (uom *UOM) Checksum() int {
	var checksum int
	for k := range uom.m {
		checksum += uom.checksumAt(k)
	}
	return checksum
}

func (uom *UOM) TransferPath(start, finish string) []string {
	path1 := uom.pathToCOM(start)
	path2 := uom.pathToCOM(finish)

	// find the length of the common suffix
	var suffixLength int
	for suffixLength = 1; path1[len(path1)-suffixLength] == path2[len(path2)-suffixLength]; suffixLength++ {
	}
	suffixLength--

	// merge path1 prefix and reverse of path2 prefix
	tp := path1[:len(path1)-suffixLength]
	for j := len(path2) - suffixLength; j >= 0; j-- {
		tp = append(tp, path2[j])
	}

	return tp
}

func buildUOM(orbits []string) *UOM {
	uom := &UOM{m: make(map[string]string)}
	for _, orbit := range orbits {
		strs := strings.Split(orbit, ")")
		uom.m[strs[1]] = strs[0]
	}
	return uom
}

func main() {
	orbits := utils.ReadLines("input.txt")
	uom := buildUOM(orbits)

	fmt.Println(uom.Checksum())
	fmt.Println(len(uom.TransferPath("YOU", "SAN")) - 1)
}
