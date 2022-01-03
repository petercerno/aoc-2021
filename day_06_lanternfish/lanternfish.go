package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const resetAge int = 6
const newbornAge int = 8

func advance(nums []int64) {
	newborns := nums[0]
	for i := 0; i < newbornAge; i++ {
		nums[i] = nums[i+1]
	}
	nums[resetAge] += newborns
	nums[newbornAge] = newborns
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("Invalid input file: Expected single line")
	}

	inp := scanner.Text()
	if scanner.Scan() {
		log.Fatal("Invalid input file: Expected EOF")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	vals := strings.Split(inp, ",")
	nums := make([]int64, newbornAge+1)
	for _, s := range vals {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		nums[n]++
	}

	for day := 0; day < 256; day++ {
		advance(nums)
	}

	var result int64
	for _, num := range nums {
		result += num
	}

	fmt.Println(result)
}
