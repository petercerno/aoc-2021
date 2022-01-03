package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func filterByColumn(lines []string, idx int, mostCommon bool) []string {
	split := make(map[rune][]string)
	for _, line := range lines {
		c := rune(line[idx])
		split[c] = append(split[c], line)
	}

	a, b := '0', '1'
	if mostCommon {
		a, b = '1', '0'
	}
	if len(split['1']) >= len(split['0']) {
		return split[a]
	} else {
		return split[b]
	}
}

func filterAll(lines []string, mostCommon bool) string {
	idx := 0
	for len(lines) > 1 && idx < len(lines[0]) {
		lines = filterByColumn(lines, idx, mostCommon)
		idx++
	}
	if len(lines) != 1 {
		log.Fatal("Expected single line. Got: ", lines)
	}

	return lines[0]
}

func binaryToInt(line string) int {
	n := 0
	z := int('0')
	for _, c := range line {
		n = 2*n + int(c) - z
	}
	return n
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	o2_line := filterAll(lines, true)
	co2_line := filterAll(lines, false)
	o2_rating := binaryToInt(o2_line)
	co2_rating := binaryToInt(co2_line)

	fmt.Println(o2_line, o2_rating, co2_line, co2_rating, o2_rating*co2_rating)
}
