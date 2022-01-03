package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	unique := map[int]bool{
		2: true,
		3: true,
		4: true,
		7: true,
	}
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		p := strings.Split(line, " | ")
		if len(p) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		out := strings.Split(p[1], " ")
		for _, s := range out {
			if unique[len(s)] {
				result++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
