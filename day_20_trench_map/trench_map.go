package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord [2]int
type image struct {
	w, h int
	m    map[coord]bool
	b    bool // background
}

var p []bool

func (im *image) get(x, y int) bool {
	if x >= 0 && x < im.w &&
		y >= 0 && y < im.h {
		return im.m[coord{x, y}]
	}
	return im.b
}

func (im *image) index(x, y int) int {
	idx := 0
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			idx = 2 * idx
			if im.get(x+i, y+j) {
				idx += 1
			}
		}
	}
	return idx
}

func (im *image) count() int {
	c := 0
	for y := 0; y < im.h; y++ {
		for x := 0; x < im.w; x++ {
			if im.get(x, y) {
				c++
			}
		}
	}
	return c
}

func enhance(im image) image {
	ime := image{
		w: im.w + 4,
		h: im.h + 4,
		m: make(map[coord]bool),
	}
	for y := -3; y <= im.h; y++ {
		for x := -3; x <= im.w; x++ {
			if p[im.index(x, y)] {
				ime.m[coord{x + 3, y + 3}] = true
			}
		}
	}
	if im.b {
		ime.b = p[0b111111111]
	} else {
		ime.b = p[0]
	}
	return ime
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	im := image{
		m: make(map[coord]bool),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if p == nil {
			p = make([]bool, len(line))
			for i := 0; i < len(line); i++ {
				p[i] = line[i] == '#'
			}
			continue
		}

		if line == "" {
			continue
		}

		if im.w == 0 {
			im.w = len(line)
		}
		if im.w != len(line) {
			log.Fatal("Invalid line: ", line)
		}

		for x := 0; x < im.w; x++ {
			if line[x] == '#' {
				im.m[coord{x, im.h}] = true
			}
		}
		im.h++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 50; i++ {
		im = enhance(im)
	}
	fmt.Println(im.count())
}
