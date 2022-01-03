package main

import (
	"container/heap"
	"fmt"
)

const (
	N = 11
	M = 5
	T = 4
)

type coord [2]int
type coords []coord
type config [N][M]int

var (
	home [T]coords
	hall coords
	move [T]map[coord]coords
	cost [T]int
)

type item struct {
	conf  *config
	cost  int
	heur  int
	index int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].cost+pq[i].heur < pq[j].cost+pq[j].heur
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil // avoid memory leak
	i.index = -1   // for safety
	*pq = old[0 : n-1]
	return i
}

func steps(c0, c1 coord) coords {
	s := make(coords, 0, N+M)
	c := c0
	s = append(s, c)
	for c != c1 {
		if c[0] != c1[0] {
			if c[1] > 0 {
				c = coord{c[0], c[1] - 1}
			} else if c[0] < c1[0] {
				c = coord{c[0] + 1, c[1]}
			} else if c[0] > c1[0] {
				c = coord{c[0] - 1, c[1]}
			}
		} else {
			if c[1] < c1[1] {
				c = coord{c[0], c[1] + 1}
			} else if c[1] > c1[1] {
				c = coord{c[0], c[1] - 1}
			}
		}
		s = append(s, c)
	}
	return s
}

func valid(c *config, s coords) bool {
	if len(s) < 2 {
		return false
	}

	c0 := s[0]
	c1 := s[len(s)-1]
	if c[c0[0]][c0[1]] == 0 {
		return false
	}

	for i := 1; i < len(s); i++ {
		if c[s[i][0]][s[i][1]] != 0 {
			return false
		}
	}

	t := c[c0[0]][c0[1]] - 1
	if c1[1] > 0 {
		for j := c1[1] + 1; j < M; j++ {
			if c[c1[0]][j] != t+1 {
				return false
			}
		}
	}

	return true
}

func validSteps(c *config) []coords {
	vs := make([]coords, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if c[i][j] > 0 {
				t := c[i][j] - 1
				c0 := coord{i, j}
				c1s := move[t][c0]
				for _, c1 := range c1s {
					s := steps(c0, c1)
					if valid(c, s) {
						vs = append(vs, s)
					}
				}
			}
		}
	}
	return vs
}

func heuristic(c *config) int {
	h := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if c[i][j] > 0 {
				t := c[i][j] - 1
				if i != 2*t+2 {
					// Type t is not at home
					l := len(steps(coord{i, j}, coord{2*t + 2, 1})) - 1
					h += cost[t] * l
				}
			}
		}
	}
	return h
}

func final(c *config) bool {
	for i := 0; i < N; i++ {
		if c[i][0] != 0 {
			return false
		}
	}

	for t := 0; t < T; t++ {
		for _, c0 := range home[t] {
			if c[c0[0]][c0[1]] != t+1 {
				return false
			}
		}
	}

	return true
}

func find(c *config) int {
	pq := make(priorityQueue, 0)
	heap.Push(&pq, &item{
		conf: c,
		cost: 0,
		heur: heuristic(c),
	})

	m := make(map[config]bool)
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		_, ok := m[*it.conf]
		if ok {
			continue
		}
		m[*it.conf] = true

		if final(it.conf) {
			return it.cost
		}

		vs := validSteps(it.conf)
		for _, s := range vs {
			var nc config = *it.conf
			c0 := s[0]
			c1 := s[len(s)-1]
			t := nc[c0[0]][c0[1]] - 1
			nc[c0[0]][c0[1]] = 0
			nc[c1[0]][c1[1]] = t + 1
			_, ok := m[nc]
			if ok {
				continue
			}

			heap.Push(&pq, &item{
				conf: &nc,
				cost: it.cost + cost[t]*(len(s)-1),
				heur: heuristic(&nc),
			})
		}
	}
	return -1
}

func main() {
	for t := 0; t < T; t++ {
		for j := 1; j < M; j++ {
			home[t] = append(home[t], coord{2*t + 2, j})
		}
	}
	hall = coords{
		{0, 0}, {1, 0},
		{3, 0}, {5, 0}, {7, 0},
		{9, 0}, {10, 0},
	}
	for t0 := 0; t0 < T; t0++ {
		move[t0] = make(map[coord]coords)
		// home -> hall
		for _, c1 := range hall {
			for i := 0; i < len(home[t0])-1; i++ {
				c0 := home[t0][i]
				move[t0][c0] = append(move[t0][c0], c1)
			}
		}
		// hall -> home
		for _, c1 := range hall {
			for _, c0 := range home[t0] {
				move[t0][c1] = append(move[t0][c1], c0)
			}
		}
		// other home -> hall / home
		for t1 := 0; t1 < T; t1++ {
			if t0 == t1 {
				continue
			}

			// other home -> hall
			for _, c0 := range home[t1] {
				for _, c1 := range hall {
					move[t0][c0] = append(move[t0][c0], c1)
				}
			}
			// other home -> home
			for _, c0 := range home[t1] {
				for _, c1 := range home[t0] {
					move[t0][c0] = append(move[t0][c0], c1)
				}
			}
		}
	}
	cost = [4]int{1, 10, 100, 1000}

	var c config
	// c[2][1] = 2
	// c[2][2] = 4
	// c[2][3] = 4
	// c[2][4] = 1
	// c[4][1] = 3
	// c[4][2] = 3
	// c[4][3] = 2
	// c[4][4] = 4
	// c[6][1] = 2
	// c[6][2] = 2
	// c[6][3] = 1
	// c[6][4] = 3
	// c[8][1] = 4
	// c[8][2] = 1
	// c[8][3] = 3
	// c[8][4] = 1

	c[2][1] = 1
	c[2][2] = 4
	c[2][3] = 4
	c[2][4] = 2
	c[4][1] = 4
	c[4][2] = 3
	c[4][3] = 2
	c[4][4] = 3
	c[6][1] = 2
	c[6][2] = 2
	c[6][3] = 1
	c[6][4] = 4
	c[8][1] = 3
	c[8][2] = 1
	c[8][3] = 3
	c[8][4] = 1

	fmt.Println(find(&c))
}
