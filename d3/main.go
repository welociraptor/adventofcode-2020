package main

import (
	"fmt"
	"log"

	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	data, err := input.Strings("d3/testdata/input")
	if err != nil {
		log.Fatal(err)
	}

	a := countTrees(data, 1, 1)
	b := countTrees(data, 3, 1)
	c := countTrees(data, 5, 1)
	d := countTrees(data, 7, 1)
	e := countTrees(data, 1, 2)

	fmt.Println("part 1:", b)
	fmt.Println("part 2:", a*b*c*d*e)
}

func countTrees(treemap []string, right, down int) int {
	rows := len(treemap)
	cols := len(treemap[0])
	trees := 0
	x := 0
	for y := down; y < rows; y += down {
		x += right
		if x >= cols {
			x -= cols
		}
		if treemap[y][x] == '#' {
			trees++
		}
	}
	return trees
}
