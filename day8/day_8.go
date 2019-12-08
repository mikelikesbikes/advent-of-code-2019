package main

import (
	"../utils"
	"fmt"
	"github.com/fatih/color"
)

type Image struct {
	height, width int
	data          string
}

func (img *Image) Layers() []string {
	var layers []string
	for i := 0; i < len(img.data); i += img.height * img.width {
		layers = append(layers, img.data[i:i+img.height*img.width])
	}
	return layers
}

func countDigits(s string) map[rune]int {
	m := make(map[rune]int)
	for _, b := range s {
		m[b]++
	}
	return m
}

func (img *Image) Checksum() int {
	layers := img.Layers()
	minDigits := countDigits(layers[0])
	for i := 1; i < len(layers); i++ {
		l := layers[i]
		if digits := countDigits(l); digits['0'] < minDigits['0'] {
			minDigits = digits
		}
	}
	return minDigits['1'] * minDigits['2']
}

func (img *Image) flatten() []int {
	layers := img.Layers()
	flat := make([]int, img.width*img.height)

	for i := 0; i < len(layers[0]); i++ {
		for _, layer := range layers {
			v := int(layer[i]) - 48
			if v != 2 {
				flat[i] = v
				break
			}
		}
	}

	return flat
}

func (img *Image) Render() {
	white := color.New(color.BgWhite).PrintfFunc()
	black := color.New(color.BgBlack).PrintfFunc()
	for i, v := range img.flatten() {
		col := i % img.width
		if col == 0 {
			fmt.Println()
		}
		if v == 0 {
			black(" ")
		} else {
			white(" ")
		}
	}
}

func main() {
	rawImage := utils.ReadLines("input.txt")[0]
	image := Image{height: 6, width: 25, data: rawImage}

	fmt.Println(image.Checksum())
	image.Render()
}
