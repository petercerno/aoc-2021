package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord [2]int
type fold [2]int

var foldPrefix string = "fold along "

func applyFold(coords []coord, f fold) []coord {
	m := make(map[coord]bool)
	result := make([]coord, 0, len(coords))
	for _, c := range coords {
		d := c
		if c[f[0]] > f[1] {
			d[f[0]] = 2*f[1] - c[f[0]]
		}
		if !m[d] {
			result = append(result, d)
			m[d] = true
		}
	}
	return result
}

func printCoords(coords []coord) {
	var max coord
	m := make(map[coord]bool)
	for _, c := range coords {
		m[c] = true
		if c[0] > max[0] {
			max[0] = c[0]
		}
		if c[1] > max[1] {
			max[1] = c[1]
		}
	}

	for y := 0; y <= max[1]; y++ {
		for x := 0; x <= max[0]; x++ {
			if m[coord{x, y}] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coords := make([]coord, 0)
	folds := make([]fold, 0)
	readCoords := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readCoords = false
			continue
		}

		if readCoords {
			c := strings.Split(line, ",")
			if len(c) != 2 {
				log.Fatal("Invalid coord line: ", line)
			}

			x, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(c[1])
			if err != nil {
				log.Fatal(err)
			}

			coords = append(coords, coord{x, y})
		} else {
			if !strings.HasPrefix(line, foldPrefix) {
				log.Fatal("Invalid fold line: ", line)
			}

			f := strings.Split(line[len(foldPrefix):], "=")
			if len(f) != 2 {
				log.Fatal("Invalid fold line: ", line)
			}

			d, err := strconv.Atoi(f[1])
			if err != nil {
				log.Fatal(err)
			}

			if f[0] == "x" {
				folds = append(folds, fold{0, d})
			} else if f[0] == "y" {
				folds = append(folds, fold{1, d})
			} else {
				log.Fatal("Invalid fold axis: ", line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, f := range folds {
		coords = applyFold(coords, f)
	}

	printCoords(coords)
}
