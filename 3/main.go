package main

import (
	"fmt"
	"os"
	"strings"
)

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func main() {
	in, _ := os.ReadFile("./3/input.txt")
	total := 0
	splt := strings.Split(string(in), "\n")
	for i, _ := range splt {
		if i%3 != 0 {
			continue
		}
		l1, l2, l3 := splt[i], splt[i+1], splt[i+2]
		ch := ""
		for _, c := range strings.Split(l1, "") {
			if strings.Contains(l2, c) && strings.Contains(l3, c) {
				ch = c
				break
			}
		}
		if strings.ToLower(ch) == ch {
			total += int(ch[0] - 64 - 26 - 6)
		} else {
			total += int(ch[0] - 64 + 26)
		}

	}
	fmt.Println(total)
}
