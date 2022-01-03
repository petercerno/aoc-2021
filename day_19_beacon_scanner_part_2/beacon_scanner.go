package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const N = 3

type coord [N]int
type coords []coord

func (a coord) addK(b coord, k int) coord {
	for i := 0; i < N; i++ {
		a[i] += b[i] * k
	}
	return a
}

func (a coord) add(b coord) coord {
	return a.addK(b, 1)
}

func (a coord) sub(b coord) coord {
	return a.addK(b, -1)
}

func (a coord) len2() int {
	d := 0
	for k := 0; k < N; k++ {
		d += a[k] * a[k]
	}
	return d
}

func (a coord) lenAbs() int {
	d := 0
	for k := 0; k < N; k++ {
		if a[k] > 0 {
			d += a[k]
		} else {
			d += -a[k]
		}
	}
	return d
}

func (a coord) less(b coord) bool {
	for i := 0; i < N; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

func (cs coords) add(a coord) coords {
	for i := 0; i < len(cs); i++ {
		cs[i] = cs[i].add(a)
	}
	return cs
}

type transform struct {
	i [N]int // coordinate indices from [0, 1, 2]
	s [N]int // coordinate signs from [-1, +1]
}

func (c coord) apply(t *transform) coord {
	return coord{
		t.s[0] * c[t.i[0]],
		t.s[1] * c[t.i[1]],
		t.s[2] * c[t.i[2]]}
}

func (cs coords) apply(t *transform) coords {
	cst := make(coords, 0, len(cs))
	for _, c := range cs {
		cst = append(cst, c.apply(t))
	}
	return cst
}

func (a coords) Len() int           { return len(a) }
func (a coords) Less(i, j int) bool { return a[i].less(a[j]) }
func (a coords) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type pmap map[coord]bool
type dmap map[int]int

var ts []*transform = make([]*transform, 0, 24)
var a []coords = make([]coords, 0)
var b []coords = make([]coords, 0)
var pm []pmap
var gd []dmap
var ld [][]dmap

func (cs coords) pmap() pmap {
	pm := make(pmap)
	for k := 0; k < len(cs); k++ {
		pm[coord{cs[k][0], cs[k][1], cs[k][2]}] = true
	}
	return pm
}

func (cs coords) glbD() dmap {
	m := make(dmap)
	for i := 0; i < len(cs)-1; i++ {
		for j := i + 1; j < len(cs); j++ {
			m[cs[i].sub(cs[j]).len2()]++
		}
	}
	return m
}

func (cs coords) locD() []dmap {
	m := make([]dmap, len(cs))
	for i := 0; i < len(cs); i++ {
		m[i] = make(dmap)
		for j := 0; j < len(cs); j++ {
			if i == j {
				continue
			}

			m[i][cs[i].sub(cs[j]).len2()]++
		}
	}
	return m
}

func intersect(m1, m2 dmap) int {
	if len(m1) > len(m2) {
		return intersect(m2, m1)
	}

	s := 0
	for d, n := range m1 {
		m, ok := m2[d]
		if !ok {
			continue
		}

		if m < n {
			s += m
		} else {
			s += n
		}
	}
	return s
}

func find(u, v int) (coords, *transform, coord) {
	if intersect(gd[u], gd[v]) < 66 {
		return nil, nil, coord{}
	}

	cs1 := a[u]
	cs2 := a[v]
	ld1 := ld[u]
	ld2 := ld[v]
	pm1 := pm[u]
	ij := make(map[[2]int]bool)
	for i := 0; i < len(cs1); i++ {
		for j := 0; j < len(cs2); j++ {
			if intersect(ld1[i], ld2[j]) >= 11 {
				ij[[2]int{i, j}] = true
			}
		}
	}
	for _, t := range ts {
		cs2t := cs2.apply(t)
		for i := 0; i < len(cs1); i++ {
			for j := 0; j < len(cs2t); j++ {
				if !ij[[2]int{i, j}] {
					continue
				}

				count := 0
				for k := 0; k < len(cs2t); k++ {
					pk := cs2t[k].sub(cs2t[j]).add(cs1[i])
					_, ok := pm1[pk]
					if ok {
						count++
					}
				}
				if count >= 12 {
					cs := make(coords, 0, len(cs1)+len(cs2t))
					for k := 0; k < len(cs1); k++ {
						cs = append(cs, cs1[k])
					}
					for k := 0; k < len(cs2t); k++ {
						pk := cs2t[k].sub(cs2t[j]).add(cs1[i])
						_, ok := pm1[pk]
						if !ok {
							cs = append(cs, pk)
						}
					}
					return cs, t, cs1[i].sub(cs2t[j])
				}
			}
		}
	}
	return nil, nil, coord{}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var c coords
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			a = append(a, c)
			b = append(b, coords{coord{}})
			continue
		} else if len(line) >= 3 && line[:3] == "---" {
			c = make(coords, 0)
			continue
		}

		v := strings.Split(line, ",")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		z, _ := strconv.Atoi(v[2])
		c = append(c, coord{x, y, z})
	}
	a = append(a, c)
	b = append(b, coords{coord{}})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i0 := 0; i0 < N; i0++ {
		for i1 := 0; i1 < N; i1++ {
			if i1 == i0 {
				continue
			}

			for s0 := +1; s0 >= -1; s0 -= 2 {
				for s1 := +1; s1 >= -1; s1 -= 2 {
					i2 := N - i0 - i1
					s2 := s0 * s1
					if i1 == (i0+2)%N {
						s2 = -s2
					}
					ts = append(ts, &transform{
						i: [N]int{i0, i1, i2},
						s: [N]int{s0, s1, s2},
					})
				}
			}
		}
	}

	pm = make([]pmap, 0, len(a))
	gd = make([]dmap, 0, len(a))
	ld = make([][]dmap, 0, len(a))
	for i := 0; i < len(a); i++ {
		pm = append(pm, a[i].pmap())
		gd = append(gd, a[i].glbD())
		ld = append(ld, a[i].locD())
	}

	u := len(a) - 2
	for u >= 0 {
		found := false
		for v := len(a) - 1; v > u; v-- {
			if a[v] == nil {
				continue
			}

			cs, t, o := find(u, v)
			if cs != nil {
				a[u], pm[u], gd[u], ld[u] = cs, cs.pmap(), cs.glbD(), cs.locD()
				a[v], pm[v], gd[v], ld[v] = nil, nil, nil, nil
				b[u] = append(b[u], b[v].apply(t).add(o)...)
				b[v] = nil
				found = true
			}
		}
		if !found {
			u--
		}
	}

	maxD := 0
	for i := 0; i < len(b[0])-1; i++ {
		for j := i + 1; j < len(b[0]); j++ {
			d := b[0][i].sub(b[0][j]).lenAbs()
			if d > maxD {
				maxD = d
			}
		}
	}
	fmt.Println(maxD)
}
