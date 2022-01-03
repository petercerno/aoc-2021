package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos := 0
	depth := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s) != 2 {
			log.Fatal("Invalid slice: ", s)
		}

		cmd := s[0]
		n, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case "forward":
			pos += n
		case "down":
			depth += n
		case "up":
			depth -= n
		default:
			log.Fatal("Invalid command: ", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pos, depth, pos*depth)
}
