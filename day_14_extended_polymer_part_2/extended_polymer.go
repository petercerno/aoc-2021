package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var r map[string]string // rules

func expand(m map[string]int64) map[string]int64 {
	var sb strings.Builder
	newM := make(map[string]int64)
	for p, n := range m {
		q, ok := r[p]
		if ok {
			sb.Reset()
			sb.WriteByte(p[0])
			sb.WriteString(q)
			newM[sb.String()] += n

			sb.Reset()
			sb.WriteString(q)
			sb.WriteByte(p[1])
			newM[sb.String()] += n
		} else {
			newM[p] += n
		}
	}
	return newM
}

func counts(m map[string]int64) (int64, int64) {
	c := make(map[byte]int64)
	for p, n := range m {
		c[p[0]] += n
		c[p[1]] += n
	}

	var max int64 = -1
	var min int64 = -1
	for b, n := range c {
		if b == '|' {
			continue
		}

		if max == -1 || n > max {
			max = n
		}
		if min == -1 || n < min {
			min = n
		}
	}

	return max / 2, min / 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r = make(map[string]string) // rules
	m := make(map[string]int64) // pair counts
	readT := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readT = false
			continue
		}

		if readT {
			t := fmt.Sprintf("|%s|", line)
			for i := 0; i < len(t)-1; i++ {
				m[t[i:i+2]]++
			}
		} else {
			p := strings.Split(line, " -> ")
			if len(p) != 2 || len(p[0]) != 2 || len(p[1]) != 1 {
				log.Fatal("Invalid rule: ", line)
			}

			r[p[0]] = p[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 40; i++ {
		m = expand(m)
	}

	a, b := counts(m)
	fmt.Println(a - b)
}
