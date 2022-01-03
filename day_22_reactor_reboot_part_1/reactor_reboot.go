package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval [2]int
type cuboid struct {
	box [3]interval
	on  bool
}
type coord [3]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func update(m map[coord]bool, c *cuboid) int {
	d := 0
	for i := max(c.box[0][0], -50); i <= min(c.box[0][1], +50); i++ {
		for j := max(c.box[1][0], -50); j <= min(c.box[1][1], +50); j++ {
			for k := max(c.box[2][0], -50); k <= min(c.box[2][1], +50); k++ {
				if m[coord{i, j, k}] != c.on {
					if c.on {
						d++
					} else {
						d--
					}
				}
				m[coord{i, j, k}] = c.on
			}
		}
	}
	return d
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	steps := make([]*cuboid, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		c := &cuboid{}
		if line[:2] == "on" {
			c.on = true
			line = line[3:]
		} else {
			line = line[4:]
		}

		ranges := strings.Split(line, ",")
		if len(ranges) != 3 {
			log.Fatal("Invalid line: ", line)
		}

		for i := 0; i < 3; i++ {
			rng := strings.Split(ranges[i][2:], "..")
			if len(rng) != 2 {
				log.Fatal("Invalid range: ", ranges[i][2:])
			}

			c.box[i][0], _ = strconv.Atoi(rng[0])
			c.box[i][1], _ = strconv.Atoi(rng[1])
		}
		steps = append(steps, c)
	}

	d := 0
	m := make(map[coord]bool)
	for _, s := range steps {
		d += update(m, s)
	}
	fmt.Println(d)
}
