package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var rows int
var cols int

type coord [2]int

func step(a [][]int) int {
	f := 0
	p := make([]coord, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			a[i][j]++
			if a[i][j] == 10 {
				p = append(p, coord{i, j})
				a[i][j] = 0
				f++
			}
		}
	}

	for len(p) > 0 {
		c := p[len(p)-1]
		p = p[0 : len(p)-1]
		for u := -1; u <= +1; u++ {
			for v := -1; v <= +1; v++ {
				i, j := c[0]+u, c[1]+v
				if i < 0 || i >= rows || j < 0 || j >= cols ||
					a[i][j] == 0 {
					continue
				}

				a[i][j]++
				if a[i][j] == 10 {
					p = append(p, coord{i, j})
					a[i][j] = 0
					f++
				}
			}
		}
	}

	return f
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a := make([][]int, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		} else if len(line) != cols {
			log.Fatal("Invalid line: ", line)
		}

		b := make([]int, cols)
		for i := 0; i < cols; i++ {
			b[i] = int(line[i] - byte('0'))
		}
		a = append(a, b)
	}
	rows = len(a)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	for {
		result++
		if step(a) == 100 {
			break
		}
	}

	fmt.Println(result)
}
