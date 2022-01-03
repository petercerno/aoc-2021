package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type line [2]point

func strToPoint(s string) point {
	p := strings.Split(s, ",")
	if len(p) != 2 {
		log.Fatal("Invalid point string: ", s)
	}

	x, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatal(err)
	}

	return point{x, y}
}

func strToLine(s string) line {
	p := strings.Split(s, " -> ")
	if len(p) != 2 {
		log.Fatal("Invalid line string: ", s)
	}

	return line{strToPoint(p[0]), strToPoint(p[1])}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, strToLine(s))
	}

	result := 0
	m := make(map[point]int)
	for _, l := range lines {
		if l[0].x == l[1].x {
			x := l[0].x
			y1, y2 := l[0].y, l[1].y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				p := point{x, y}
				m[p]++
				if m[p] == 2 {
					result++
				}
			}
		} else if l[0].y == l[1].y {
			y := l[0].y
			x1, x2 := l[0].x, l[1].x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				p := point{x, y}
				m[p]++
				if m[p] == 2 {
					result++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
