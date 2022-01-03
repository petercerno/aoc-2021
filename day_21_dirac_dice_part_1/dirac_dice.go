package main

import "fmt"

func main() {
	p := [2]int{2, 6}
	s := [2]int{}
	d := 0
	for s[0] < 1000 && s[1] < 1000 {
		for i := 0; i < 2; i++ {
			p[i] = (p[i] + 3*d + 6) % 10
			s[i] += p[i] + 1
			d += 3
			if s[i] >= 1000 {
				fmt.Println(s[1-i] * d)
				break
			}
		}
	}
}
