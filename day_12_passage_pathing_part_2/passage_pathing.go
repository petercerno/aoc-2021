package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var e map[string][]string
var m map[string]int
var p []string
var r bool
var c int

func add(a, b string) {
	if _, ok := e[a]; !ok {
		e[a] = make([]string, 0)
	}
	if _, ok := e[b]; !ok {
		e[b] = make([]string, 0)
	}
	e[a] = append(e[a], b)
	e[b] = append(e[b], a)
}

func big(s string) bool {
	for _, r := range s {
		return unicode.IsUpper(r)
	}
	return false
}

func explore(u string) {
	m[u]++
	p = append(p, u)
	defer func(u string) {
		m[u]--
		p = p[0 : len(p)-1]
	}(u)

	if u == "end" {
		c++
		return
	}

	for _, v := range e[u] {
		if m[v] == 0 || big(v) {
			explore(v)
		} else if m[v] == 1 && v != "start" && !r {
			r = true
			explore(v)
			r = false
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	e = make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		add(s[0], s[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m = make(map[string]int)
	p = make([]string, 0)
	explore("start")

	fmt.Println(c)
}
