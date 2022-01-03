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
var m map[string]bool
var p []string
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
	m[u] = true
	p = append(p, u)
	defer func(u string) {
		m[u] = false
		p = p[0 : len(p)-1]
	}(u)

	if u == "end" {
		c++
		return
	}

	for _, v := range e[u] {
		if !m[v] || big(v) {
			explore(v)
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

	m = make(map[string]bool)
	p = make([]string, 0)
	explore("start")

	fmt.Println(c)
}
