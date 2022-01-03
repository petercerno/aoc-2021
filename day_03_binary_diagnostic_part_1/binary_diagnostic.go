package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n := 0
	var a []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if a == nil {
			a = make([]int, len(s))
		}
		for i, c := range s {
			a[i] += int(c) - int('0')
		}
		n++
	}

	gamma := 0
	eps := 0
	for _, x := range a {
		d := 0
		if x > n-x {
			d = 1
		}
		gamma = 2*gamma + d
		eps = 2*eps + 1 - d
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(a, gamma, eps, gamma*eps)
}
