package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.ReadFile("./6/input.txt")
	l := 14
	for i, _ := range string(in) {
		r := string(in)[int(math.Max(float64(i-l), 0)):i]
		m := make(map[int32]bool)
		for _, l := range r {
			m[l] = true
		}
		fmt.Println(m)
		if len(m) == l {
			fmt.Println(i)
			break
		}
	}
}
