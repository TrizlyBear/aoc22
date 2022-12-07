package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("./4/input.txt")
	total := 0
	for _, row := range strings.Split(string(in), "\n") {
		l, r := strings.Split(row, ",")[0], strings.Split(row, ",")[1]
		lbl, lbr := atoi(strings.Split(l, "-")[0]), atoi(strings.Split(l, "-")[1])
		rbl, rbr := atoi(strings.Split(r, "-")[0]), atoi(strings.Split(r, "-")[1])
		if (lbr >= rbl && lbl <= rbr) || (rbr >= lbl && rbl <= lbr) {
			total += 1
		}
	}
	fmt.Println(total)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
