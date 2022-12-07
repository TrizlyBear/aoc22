package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("./5/input.txt")
	stacks, moves := strings.Split(string(in), "\n\n")[0], strings.Split(string(in), "\n\n")[1]
	stack := make([][]string, len(strings.Split(stacks, "\n")))
	for i, l := range strings.Split(stacks, "\n") {
		if i == len(strings.Split(stacks, "\n"))-1 {
			break
		}
		for ic, c := range l {
			if (ic-1)%4 == 0 && string(c) != " " {
				//fmt.Println((ic-1)/4, string(c))
				stack[(ic-1)/4] = append([]string{string(c)}, stack[(ic-1)/4]...)
			}
		}
	}
	fmt.Println(stack)
	_ = stack

	for _, l := range strings.Split(moves, "\n") {
		split := strings.Split(l, " ")
		amt := atoi(split[1])
		from := atoi(split[3])
		to := atoi(split[5])
		mov := stack[from-1][len(stack[from-1])-amt:]
		stack[from-1] = stack[from-1][:len(stack[from-1])-amt]
		stack[to-1] = append(stack[to-1], mov...)
		fmt.Println("New: ", stack)
	}
	fmt.Println(stack)
	for _, l := range stack {
		if len(l) != 0 {
			fmt.Print(l[len(l)-1])
		}
	}
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		temp := s[len(s)-i-1]
		s[len(s)-i-1] = s[i]
		s[i] = temp
	}
	return s
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
