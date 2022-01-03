package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Sliding window
type window struct {
	len  int   // Length of the window
	data []int // Elements in the window
	idx  int   // Index of the oldest element in the full window
	sum  int   // Sum of all elements in the window
}

func (w *window) add(e int) {
	if !w.full() {
		w.data = append(w.data, e)
		w.sum += e
	} else {
		w.sum -= w.data[w.idx]
		w.data[w.idx] = e
		w.sum += e
	}
	w.idx = (w.idx + 1) % w.len
}

func (w *window) full() bool {
	return len(w.data) == w.len
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := &window{len: 3} // Current window
	prev := -1           // Sum of elements in the previous window
	res := 0             // Total number of increases
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		e, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		w.add(e)
		if w.full() {
			if prev >= 0 && w.sum > prev {
				res++
			}
			prev = w.sum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
