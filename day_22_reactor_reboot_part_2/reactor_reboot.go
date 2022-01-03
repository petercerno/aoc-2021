package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	at "github.com/golang-collections/go-datastructures/augmentedtree"
)

const numDim = 3

type dimension struct {
	low, high int
}

type interval struct {
	id  uint64
	on  bool
	dim []*dimension
}

func (iv *interval) Volume() int {
	vol := 1
	for i := 0; i < len(iv.dim); i++ {
		vol *= iv.dim[i].high - iv.dim[i].low
	}
	return vol
}

func (iv *interval) LowAtDimension(d uint64) int64 {
	return int64(iv.dim[d-1].low)
}

func (iv *interval) HighAtDimension(d uint64) int64 {
	return int64(iv.dim[d-1].high)
}

func (iv1 *interval) OverlapsAtDimension(iv2 at.Interval, d uint64) bool {
	return iv1.HighAtDimension(d) > iv2.LowAtDimension(d) &&
		iv1.LowAtDimension(d) < iv2.HighAtDimension(d)
}

func (iv *interval) ID() uint64 {
	return iv.id
}

var id uint64 // global interval id counter

type splitStruct struct {
	split  [numDim][]int
	index  [numDim]int
	output []at.Interval
}

// split interval ov into smaller sub-intervals disjoint with iv
func split(ov *interval, iv *interval) []at.Interval {
	ss := &splitStruct{}
	for d := 0; d < numDim; d++ {
		ss.split[d] = []int{
			ov.dim[d].low, ov.dim[d].high,
			iv.dim[d].low, iv.dim[d].high}
		sort.Ints(ss.split[d])
	}
	splitImpl(ov, iv, 0, ss)
	return ss.output
}

func splitImpl(ov *interval, iv *interval, d int, ss *splitStruct) {
	if d < numDim {
		// select interval at dimension d
		for i := 0; i < len(ss.split[d])-1; i++ {
			if ss.split[d][i] < ov.dim[d].low ||
				ss.split[d][i] >= ov.dim[d].high ||
				ss.split[d][i] == ss.split[d][i+1] {
				continue
			}

			ss.index[d] = i
			splitImpl(ov, iv, d+1, ss)
		}
	} else if d == numDim {
		// all intervals have been selected
		id++
		out := &interval{
			id:  id,
			on:  ov.on,
			dim: make([]*dimension, numDim),
		}
		disjoint := false
		for i := 0; i < numDim; i++ {
			out.dim[i] = &dimension{
				low:  ss.split[i][ss.index[i]],
				high: ss.split[i][ss.index[i]+1],
			}
			if iv.dim[i].high <= out.dim[i].low ||
				out.dim[i].high <= iv.dim[i].low {
				// interval out is disjoint with iv
				disjoint = true
			}
		}
		if disjoint {
			ss.output = append(ss.output, out)
		}
	}
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cuboids := make([]*interval, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id++
		iv := &interval{id: id}
		if line[:2] == "on" {
			iv.on = true
			line = line[3:]
		} else {
			line = line[4:]
		}
		ranges := strings.Split(line, ",")
		iv.dim = make([]*dimension, numDim)
		for i := 0; i < numDim; i++ {
			rng := strings.Split(ranges[i][2:], "..")
			l, _ := strconv.Atoi(rng[0])
			h, _ := strconv.Atoi(rng[1])
			iv.dim[i] = &dimension{low: l, high: h + 1}
		}
		cuboids = append(cuboids, iv)
	}

	tree := at.New(numDim)
	for _, iv := range cuboids {
		oivs := tree.Query(iv)         // intervals overlapping with iv
		sivs := make([]at.Interval, 0) // splitted overlapping intervals (disjoint iv)
		for _, ov_ := range oivs {
			ov := ov_.(*interval)
			sivs = append(sivs, split(ov, iv)...)
		}
		tree.Delete(oivs...)
		tree.Add(sivs...)
		tree.Add(iv)
	}

	all := tree.Query(&interval{
		dim: []*dimension{
			{low: MinInt, high: MaxInt},
			{low: MinInt, high: MaxInt},
			{low: MinInt, high: MaxInt}}})
	on := 0
	for _, iv_ := range all {
		iv := iv_.(*interval)
		if iv.on {
			on += iv.Volume()
		}
	}
	fmt.Println(on)
}
