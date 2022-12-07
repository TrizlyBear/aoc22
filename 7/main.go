package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name   string
	files  map[string]int
	dirs   map[string]*dir
	parent *dir
}

func main() {
	in, _ := os.ReadFile("./7/input.txt")
	d := &dir{"/", make(map[string]int), make(map[string]*dir), nil}
	point := d
	skip := 0
	split := strings.Split(string(in), "\n")
	for i, l := range split[1:] {
		//fmt.Println(l)
		if skip > 0 {
			skip--
			continue
		}
		if l == "$ ls" {
			for _, e := range split[i+2:] {
				fmt.Println(e)
				if strings.HasPrefix(e, "$") {
					break
				} else if strings.HasPrefix(e, "dir") {
					fmt.Println(strings.Split(e, " ")[1])
					point.dirs[strings.Split(e, " ")[1]] = &dir{name: strings.Split(e, " ")[1], files: make(map[string]int), parent: point, dirs: map[string]*dir{}}
					skip++
				} else {
					point.files[strings.Split(e, " ")[1]] = atoi(strings.Split(e, " ")[0])
					skip++
				}
			}
		} else if l == "$ cd .." {
			fmt.Println("cd up")
			point = point.parent
		} else if strings.HasPrefix(l, "$ cd") {
			d := strings.Split(l, " ")[2]
			point = point.dirs[d]
		}
	}

	smallest := 10000000000000
	d.traverse(func(di *dir) {
		if d.size()-di.size() <= 40000000 && di.size() < smallest {
			smallest = di.size()
		}
	})
	fmt.Println(smallest)
}

func (d *dir) size() int {
	total := 0
	for _, child := range d.dirs {
		total += child.size()
	}
	for _, f := range d.files {
		total += f
	}
	return total
}

func (d *dir) traverse(f func(*dir)) {
	for _, child := range d.dirs {
		child.traverse(f)
	}
	f(d)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
