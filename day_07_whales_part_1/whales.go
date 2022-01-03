package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	for _, s := range vals {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, n)
	}

	sort.Ints(a)
	med := a[(len(a)-1)/2]
	sum := 0
	for _, n := range a {
		if n < med {
			sum += med - n
		} else {
			sum += n - med
		}
	}
	fmt.Println(sum)
}
