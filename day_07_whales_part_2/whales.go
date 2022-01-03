package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("Invalid input file: Expected single line")
	}

	inp := scanner.Text()
	if scanner.Scan() {
		log.Fatal("Invalid input file: Expected EOF")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	vals := strings.Split(inp, ",")
	a := make([]int, 0, len(vals))
	sum := 0
	for _, s := range vals {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, n)
		sum += n
	}
	mean := sum / len(a)

	best_cost := -1
	for i := -3; i <= 3; i++ {
		cost := 0
		for _, n := range a {
			dist := n - (mean + i)
			if dist < 0 {
				dist = -dist
			}
			cost += dist * (dist + 1) / 2
		}
		if best_cost == -1 || cost < best_cost {
			best_cost = cost
		}
	}
	fmt.Println(best_cost)
}
