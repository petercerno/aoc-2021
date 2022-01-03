package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var m map[byte]byte

func illegal(s string) int {
	a := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		c, ok := m[b]
		if ok {
			if len(a) == 0 || a[len(a)-1] != c {
				return i
			}

			a = a[0 : len(a)-1]
		} else {
			a = append(a, b)
		}
	}
	return -1
}

func score(b byte) int {
	switch b {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i := illegal(line)
		if i >= 0 {
			result += score(line[i])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
