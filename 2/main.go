package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.ReadFile("./2/input.txt")
	total := 0
	for _, s := range strings.Split(string(in), "\n") {
		y, o := strings.Split(s, " ")[1], strings.Split(s, " ")[0]

		if y == "Y" {
			total += 3
		} else if y == "Z" {
			total += 6
		}
		// scissors
		if (o == "A" && y == "X") || (y == "Y" && o == "C") || (y == "Z" && o == "B") {
			total += 3
		} else if (y == "Y" && o == "B") || (y == "Z" && o == "A") || (y == "X" && o == "C") {
			total += 2
		} else {
			total += 1
		}
	}
	fmt.Println(total)
}
