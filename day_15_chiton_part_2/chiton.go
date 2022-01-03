package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type coord [2]int
type item struct {
	pos   coord
	risk  int
	index int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].risk < pq[j].risk
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

var m [][]int
var rows int
var cols int

const N int = 5

func findPath() int {
	offsets := []coord{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}
	r := make(map[coord]int)
	pq := make(priorityQueue, 0, 4*N*(rows+cols))
	heap.Push(&pq, &item{pos: coord{0, 0}, risk: 0})
	for len(pq) > 0 {
		i := heap.Pop(&pq).(*item)
		if _, ok := r[i.pos]; ok {
			continue
		}

		r[i.pos] = i.risk
		if i.pos[0] == N*rows-1 && i.pos[1] == N*cols-1 {
			return i.risk
		}

		for _, o := range offsets {
			c := coord{i.pos[0] + o[0], i.pos[1] + o[1]}
			if c[0] < 0 || c[0] >= N*rows ||
				c[1] < 0 || c[1] >= N*cols {
				continue
			}
			if _, ok := r[c]; ok {
				continue
			}

			rowR := c[0] / rows
			colR := c[1] / cols
			risk := m[c[0]%rows][c[1]%cols] + rowR + colR
			risk = (risk-1)%9 + 1

			heap.Push(&pq, &item{pos: c, risk: i.risk + risk})
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		}
		if len(line) != cols {
			log.Fatal("Invalid line: ", line)
		}

		a := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			a[i] = int(line[i]) - int('0')
		}
		m = append(m, a)
	}
	rows = len(m)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(findPath())
}
