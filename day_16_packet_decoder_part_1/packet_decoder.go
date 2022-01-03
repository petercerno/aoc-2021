package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	literal int = 4
)

type packet struct {
	version int
	typeId  int
	value   int
	subP    []*packet
}

func bin2int(b []byte) int {
	n := 0
	for i := 0; i < len(b); i++ {
		n = 2*n + int(b[i])
	}
	return n
}

func readValue(b []byte) (int, []byte) {
	f := true
	n := 0
	for f {
		f = b[0] == 1
		n = 16*n + bin2int(b[1:5])
		b = b[5:]
	}
	return n, b
}

func readPacket(b []byte) (*packet, []byte) {
	p := &packet{
		version: bin2int(b[:3]),
		typeId:  bin2int(b[3:6]),
	}
	b = b[6:]
	if p.typeId == literal {
		p.value, b = readValue(b)
	} else {
		p.subP = make([]*packet, 0)
		if b[0] == 0 {
			numB := bin2int(b[1:16])
			b = b[16:]
			c := b[:numB]
			for len(c) > 0 {
				q, nc := readPacket(c)
				p.subP = append(p.subP, q)
				c = nc
			}
			b = b[numB:]
		} else {
			numP := bin2int(b[1:12])
			b = b[12:]
			for i := 0; i < numP; i++ {
				q, nb := readPacket(b)
				p.subP = append(p.subP, q)
				b = nb
			}
		}
	}
	return p, b
}

func sumVersion(p *packet) int {
	result := p.version
	for _, q := range p.subP {
		result += sumVersion(q)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var h string
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		h = scanner.Text()
	}
	if scanner.Scan() {
		log.Fatal("Single line expected")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := map[byte][]byte{
		'0': {0, 0, 0, 0}, '1': {0, 0, 0, 1}, '2': {0, 0, 1, 0}, '3': {0, 0, 1, 1},
		'4': {0, 1, 0, 0}, '5': {0, 1, 0, 1}, '6': {0, 1, 1, 0}, '7': {0, 1, 1, 1},
		'8': {1, 0, 0, 0}, '9': {1, 0, 0, 1}, 'A': {1, 0, 1, 0}, 'B': {1, 0, 1, 1},
		'C': {1, 1, 0, 0}, 'D': {1, 1, 0, 1}, 'E': {1, 1, 1, 0}, 'F': {1, 1, 1, 1},
	}
	b := make([]byte, 0, 4*len(h))
	for i := 0; i < len(h); i++ {
		b = append(b, m[h[i]]...)
	}

	p, _ := readPacket(b)
	fmt.Println(sumVersion(p))
}
