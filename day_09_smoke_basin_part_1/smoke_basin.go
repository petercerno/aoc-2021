package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var rows, cols int
var h [][]int = make([][]int, 0)

func lowPoint(i, j int) bool {
	return (i == 0 || h[i][j] < h[i-1][j]) &&
		(i == rows-1 || h[i][j] < h[i+1][j]) &&
		(j == 0 || h[i][j] < h[i][j-1]) &&
		(j == cols-1 || h[i][j] < h[i][j+1])
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

	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if lowPoint(i, j) {
				result += h[i][j] + 1
			}
		}
	}

	fmt.Println(result)
}
