package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var rows, cols int
var h [][]int = make([][]int, 0)
var x [][]bool

func lowPoint(i, j int) bool {
	return (i == 0 || h[i][j] < h[i-1][j]) &&
		(i == rows-1 || h[i][j] < h[i+1][j]) &&
		(j == 0 || h[i][j] < h[i][j-1]) &&
		(j == cols-1 || h[i][j] < h[i][j+1])
}

func explore(i, j int) int {
	if i < 0 || i >= rows || j < 0 || j >= cols ||
		x[i][j] || h[i][j] == 9 {
		return 0
	}
	x[i][j] = true
	return 1 +
		explore(i-1, j) + explore(i+1, j) +
		explore(i, j-1) + explore(i, j+1)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		}
		a := make([]int, len(line))
		for i, c := range line {
			a[i] = int(c) - int('0')
		}
		h = append(h, a)
	}
	rows = len(h)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	x = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		x[i] = make([]bool, cols)
	}

	r := make([]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if lowPoint(i, j) {
				r = append(r, explore(i, j))
			}
		}
	}

	sort.Ints(r)
	r = r[len(r)-3:]

	fmt.Println(r[0] * r[1] * r[2])
}
