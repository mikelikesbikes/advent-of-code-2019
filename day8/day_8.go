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

func (img *Image) Render() {
	white := color.New(color.BgWhite).PrintfFunc()
	black := color.New(color.BgBlack).PrintfFunc()

	layers := img.Layers()
	for i := 0; i < len(layers[0]); i++ {
		col := i % img.width
		if col == 0 {
			fmt.Println()
		}
		for _, layer := range layers {
			v := layer[i] - 48
			if v != 2 {
				if v == 0 {
					black(" ")
				} else {
					white(" ")
				}
				break
			}
		}
	}
}

func main() {
	rawImage := utils.ReadLines("input.txt")[0]
	image := Image{height: 6, width: 25, data: rawImage}
	layers := image.Layers()
	minDigits := countDigits(layers[0])
	for i := 1; i < len(layers); i++ {
		l := layers[i]
		if digits := countDigits(l); digits['0'] < minDigits['0'] {
			minDigits = digits
		}
	}
	fmt.Println(minDigits['1'] * minDigits['2'])

	image.Render()
}
