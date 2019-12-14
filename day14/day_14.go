package main

import (
	"../utils"
	"fmt"
	"math"
	"strings"
)

type Chemical struct {
	name   string
	amount int
}

type Reaction struct {
	product Chemical
	inputs  []Chemical
}

type Factory struct {
	reactions map[string]Reaction
}

func ParseChemical(s string) Chemical {
	strs := strings.Split(s, " ")
	return Chemical{strs[1], utils.Atoi(strs[0])}
}

func ParseReaction(s string) Reaction {
	strs := strings.Split(s, " => ")
	product := ParseChemical(strs[1])
	var inputs []Chemical
	for _, str := range strings.Split(strs[0], ", ") {
		inputs = append(inputs, ParseChemical(str))
	}
	return Reaction{product: product, inputs: inputs}
}

func ParseFactory(lines []string) *Factory {
	reactions := make(map[string]Reaction, len(lines))
	for _, line := range lines {
		reaction := ParseReaction(line)
		reactions[reaction.product.name] = reaction
	}
	return &Factory{reactions}
}

func reductionsDone(chemicals map[string]int) bool {
	for k, v := range chemicals {
		if k != "ORE" && v > 0 {
			return false
		}
	}
	return true
}

func (f *Factory) rawMaterialsFor2(name string, amount int) int {
	materials := map[string]int{name: amount}
	for !reductionsDone(materials) {
		fmt.Println(materials)
		reduced := make(map[string]int)
		for k, amtNeeded := range materials {
			if amtNeeded <= 0 || k == "ORE" {
				reduced[k] += amtNeeded
				continue
			}
			reaction := f.reactions[k]
			mul := int(math.Ceil(float64(amtNeeded) / float64(reaction.product.amount)))
			for _, chem := range reaction.inputs {
				reduced[chem.name] += chem.amount * mul
			}
			reduced[k] = amtNeeded - (reaction.product.amount * mul)
		}
		materials = reduced
	}
	fmt.Println(materials)
	return materials["ORE"]
}

func (f *Factory) rawMaterialsFor(name string, amount int) int {
	materials := map[string]int{name: amount}
	for !reductionsDone(materials) {
		var k string
		var v int
		for k, v = range materials {
			if k != "ORE" && v > 0 {
				break
			}
		}
		r := f.reactions[k]
		mul := int(math.Ceil(float64(v) / float64(r.product.amount)))
		materials[r.product.name] -= (r.product.amount * mul)
		for _, chem := range r.inputs {
			materials[chem.name] += (chem.amount * mul)
		}
	}
	return materials["ORE"]
}

func (f *Factory) maxFuelFor(name string, amount int) int {
	min, max := 0, amount
	for max > min {
		mid := min + ((max - min + 1) / 2)
		oreNeeded := f.rawMaterialsFor("FUEL", mid)
		if oreNeeded > amount {
			max = mid - 1
		} else {
			min = mid
		}
	}
	return min
}

func main() {
	lines := utils.ReadLines("input.txt")

	factory := ParseFactory(lines)
	fmt.Println(factory.rawMaterialsFor("FUEL", 1))

	fmt.Println(factory.maxFuelFor("ORE", 1000000000000))
}
