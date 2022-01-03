package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type coord struct {
	row, col int
}

type board struct {
	a         [][]int
	m         map[int]coord
	sum       int
	sumMarked int
	rowMarked []int
	colMarked []int
}

func strsToInts(vals []string) []int {
	res := make([]int, 0, len(vals))
	for _, val := range vals {
		if val == "" {
			continue
		}

		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, n)
	}
	return res
}

func initBoard(boardLines []string) *board {
	if len(boardLines) != boardSize {
		log.Fatal("Invalid number of rows:\n", strings.Join(boardLines, "\n"))
	}

	b := &board{
		a:         make([][]int, boardSize),
		m:         make(map[int]coord),
		rowMarked: make([]int, boardSize),
		colMarked: make([]int, boardSize),
	}
	for row, line := range boardLines {
		cols := strsToInts(strings.Split(line, " "))
		if len(cols) != boardSize {
			log.Fatal("Invalid number of columns: ", cols)
		}

		b.a[row] = cols
		for col, n := range cols {
			b.m[n] = coord{row, col}
			b.sum += n
		}
	}
	return b
}

func (b *board) mark(n int) bool {
	c, ok := b.m[n]
	if ok {
		b.sumMarked += n
		b.rowMarked[c.row]++
		b.colMarked[c.col]++
		if b.rowMarked[c.row] == boardSize || b.colMarked[c.col] == boardSize {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 0
	var nums []int
	boardLines := make([]string, 0, boardSize)
	boards := make([]*board, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			nums = strsToInts(strings.Split(line, ","))
		} else {
			if line == "" {
				if len(boardLines) == boardSize {
					boards = append(boards, initBoard(boardLines))
					boardLines = make([]string, 0, boardSize)
				}
			} else {
				boardLines = append(boardLines, line)
			}
		}
		i++
	}
	boards = append(boards, initBoard(boardLines))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	for _, n := range nums {
		for _, b := range boards {
			if b.mark(n) {
				unmarked := b.sum - b.sumMarked
				result = unmarked * n
				break
			}
		}
		if result > 0 {
			break
		}
	}
	fmt.Println(result)
}
