package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

type permState struct {
	p []int   // Current permutation
	q []bool  // Is the number taken
	r [][]int // All permutations
}

// perms returns a slice of all permutations of n number: 0, 1, ..., n-1.
func perms(n int) [][]int {
	p := &permState{
		p: make([]int, n),
		q: make([]bool, n),
		r: make([][]int, 0, factorial(n)),
	}
	permsImpl(n, 0, p)
	return p.r
}

func permsImpl(n, i int, p *permState) {
	if i == n {
		out := make([]int, n)
		copy(out, p.p)
		p.r = append(p.r, out)
		return
	}
	for k := 0; k < n; k++ {
		if !p.q[k] {
			p.q[k] = true
			p.p[i] = k
			permsImpl(n, i+1, p)
			p.q[k] = false
		}
	}
}

// convert the given signal pattern s to a digit using the given permutation p
func convert(s string, p []int, m map[string]int) (int, bool) {
	t := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		t[i] = '0'
	}
	for i := 0; i < len(s); i++ {
		t[p[s[i]-byte('a')]] = '1'
	}
	d, ok := m[string(t)]
	return d, ok
}

// check if all signal patterns can be converted to a digit using the given permutation p
func check(a []string, p []int, m map[string]int) bool {
	for _, s := range a {
		_, ok := convert(s, p, m)
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	// segment patterns to digits
	m := map[string]int{
		"1110111": 0,
		"0010010": 1,
		"1011101": 2,
		"1011011": 3,
		"0111010": 4,
		"1101011": 5,
		"1101111": 6,
		"1010010": 7,
		"1111111": 8,
		"1111011": 9,
	}
	ps := perms(7) // all permutations
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		p := strings.Split(line, " | ")
		if len(p) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		inp := strings.Split(p[0], " ")
		out := strings.Split(p[1], " ")
		for _, p := range ps {
			if check(inp, p, m) {
				num := 0
				for _, s := range out {
					d, ok := convert(s, p, m)
					if !ok {
						log.Fatal("Cannot decode: ", s)
					}

					num = 10*num + d
				}
				result += num
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
