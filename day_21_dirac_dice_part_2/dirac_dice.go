package main

import "fmt"

type coord struct {
	p [2]int // player positions (0-based)
	s [2]int // player scores
	n int    // total number of moves
}

var r []int // r[i] == number of ways to roll i+3

func evolve(m0 map[coord]int) (map[coord]int, [2]int) {
	var wins [2]int
	m1 := make(map[coord]int)
	for c0, k0 := range m0 {
		if c0.s[0] >= 21 {
			wins[0] += k0
			continue
		}

		if c0.s[1] >= 21 {
			wins[1] += k0
			continue
		}

		i := c0.n % 2 // player that moves
		for j := 0; j < 7; j++ {
			var c1 coord = c0
			c1.p[i] = (c1.p[i] + j + 3) % 10
			c1.s[i] += c1.p[i] + 1
			c1.n++
			m1[c1] += k0 * r[j]
		}
	}
	return m1, wins
}

func main() {
	r = []int{1, 3, 6, 7, 6, 3, 1}

	var wins [2]int
	m := make(map[coord]int)
	m[coord{p: [2]int{2, 6}}] = 1

	for len(m) > 0 {
		m_, w_ := evolve(m)
		m = m_
		wins[0] += w_[0]
		wins[1] += w_[1]
	}

	if wins[0] > wins[1] {
		fmt.Println(wins[0])
	} else {
		fmt.Println(wins[1])
	}
}
