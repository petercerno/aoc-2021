package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var m map[byte]byte
var v map[byte]int

func autocomplete(s string) int {
	a := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		c, ok := m[b]
		if ok {
			if len(a) == 0 || a[len(a)-1] != c {
				return -1
			}

			a = a[0 : len(a)-1]
		} else {
			a = append(a, b)
		}
	}

	score := 0
	for i := len(a) - 1; i >= 0; i-- {
		score = 5*score + v[a[i]]
	}
	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	v = map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	scores := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		score := autocomplete(line)
		if score == -1 {
			continue
		}

		scores = append(scores, score)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(scores)
	fmt.Println(scores[(len(scores)-1)/2])
}
