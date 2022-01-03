package main

import "fmt"

const n = 14

var (
	w    [n]int
	divZ [n]int = [n]int{1, 1, 1, 1, 26, 1, 1, 26, 1, 26, 26, 26, 26, 26}
	addZ [n]int = [n]int{14, 13, 15, 13, -2, 10, 13, -15, 11, -9, -9, -7, -4, -6}
	addW [n]int = [n]int{0, 12, 14, 0, 3, 15, 11, 12, 1, 12, 3, 10, 14, 12}
	maxZ [n]int
)

func find(i, z int) bool {
	if i == n {
		return z == 0
	}

	if z > maxZ[i] || z < -maxZ[i] {
		return false
	}

	for d := 1; d <= 9; d++ {
		w[i] = d
		nz := z / divZ[i]
		if (z%26)+addZ[i] != w[i] {
			nz = 26*nz + w[i] + addW[i]
		}
		if find(i+1, nz) {
			return true
		}
	}

	return false
}

func main() {
	maxZ[n-1] = divZ[n-1]
	for i := n - 2; i >= 0; i-- {
		maxZ[i] = divZ[i] * maxZ[i+1]
	}
	find(0, 0)
	for i := 0; i < len(w); i++ {
		fmt.Print(w[i])
	}
	fmt.Println()
}
