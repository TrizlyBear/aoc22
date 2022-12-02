package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	all := [][]int{{}}
	in, _ := os.ReadFile("./1/input.txt")
	for _, s := range strings.Split(string(in), "\n") {
		if s == "" {
			all = append(all, []int{})
			continue
		}
		i, _ := strconv.Atoi(s)
		all[len(all)-1] = append(all[len(all)-1], i)
	}

	sums := []int{}
	for _, l := range all {
		sums = append(sums, sum(l...))
		//if sum(l...) > largest {
		//	largest = sum(l...)
		//}
	}
	sort.Ints(sums)
	fmt.Println(sum(sums[len(sums)-3:]...))
}

func sum(i ...int) int {
	o := 0
	for _, in := range i {
		o += in
	}
	return o
}
