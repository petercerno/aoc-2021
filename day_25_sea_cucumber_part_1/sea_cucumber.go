package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	right = iota
	down
)

type coord [2]int
type cuc struct {
	p coord // position
	t int   // type (right / down)
}
type state struct {
	M, N int              // size of the map
	m    [][]*cuc         // map of all cucumbers
	f    [2]map[*cuc]bool // free cucumbers (for a given type)
}

func (s *state) Cuc(p coord) *cuc {
	return s.m[p[0]][p[1]]
}

func (s *state) Shift(p coord, di, dj int) coord {
	return coord{(p[0] + di + s.M) % s.M, (p[1] + dj + s.N) % s.N}
}

func (s *state) Ahead(c *cuc) coord {
	switch c.t {
	case right:
		return s.Shift(c.p, 0, 1)
	case down:
		return s.Shift(c.p, 1, 0)
	}
	return coord{-1, -1}
}

func (s *state) IsFree(c *cuc) bool {
	return s.Cuc(s.Ahead(c)) == nil
}

func (s *state) Init() {
	for i := 0; i < s.M; i++ {
		for j := 0; j < s.N; j++ {
			c := s.m[i][j]
			if c != nil && s.IsFree(c) {
				s.f[c.t][c] = true
			}
		}
	}
}

func (s *state) Move(t int) {
	p0s := make([]coord, 0, len(s.f[t])) // original positions
	p1s := make([]coord, 0, len(s.f[t])) // new positions
	for c := range s.f[t] {
		p0 := c.p
		p1 := s.Ahead(c)
		p0s = append(p0s, p0)
		p1s = append(p1s, p1)
		c.p = p1
		s.m[p0[0]][p0[1]] = nil
		s.m[p1[0]][p1[1]] = c
	}
	for _, p0 := range p0s {
		// original positions are not occupied
		cl := s.Cuc(s.Shift(p0, 0, -1))
		if cl != nil && cl.t == right {
			s.f[0][cl] = true
		}
		cu := s.Cuc(s.Shift(p0, -1, 0))
		if cu != nil && cu.t == down {
			s.f[1][cu] = true
		}
	}
	for _, p1 := range p1s {
		// new positions are occupied
		c1 := s.Cuc(p1)
		if !s.IsFree(c1) {
			delete(s.f[t], c1)
		}
		cl := s.Cuc(s.Shift(p1, 0, -1))
		if cl != nil && cl.t == right {
			delete(s.f[0], cl)
		}
		cu := s.Cuc(s.Shift(p1, -1, 0))
		if cu != nil && cu.t == down {
			delete(s.f[1], cu)
		}
	}
}

func (s *state) Evolve() int {
	n := 0
	for len(s.f[right]) > 0 ||
		len(s.f[down]) > 0 {
		s.Move(right)
		s.Move(down)
		n++
	}
	return n
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := &state{}
	s.f[0] = make(map[*cuc]bool)
	s.f[1] = make(map[*cuc]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if s.N == 0 {
			s.N = len(line)
		}
		a := make([]*cuc, s.N)
		for i := 0; i < s.N; i++ {
			if line[i] == '>' {
				a[i] = &cuc{
					p: coord{s.M, i},
					t: right,
				}
			} else if line[i] == 'v' {
				a[i] = &cuc{
					p: coord{s.M, i},
					t: down,
				}
			}
		}
		s.m = append(s.m, a)
		s.M++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s.Init()
	fmt.Println(s.Evolve() + 1)
}
