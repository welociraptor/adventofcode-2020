package main

import (
	"fmt"

	"github.com/welociraptor/adventofcode-2020/v2/d7/rules"
	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	rs, err := input.Strings("d7/testdata/input", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	bags := make([]rules.Bag, 0)

	for _, r := range rs {
		bags = append(bags, rules.ParseSingle(r))
	}

	color := "shiny gold"
	canHold := make(map[string]bool)

	for _, h := range holders(bags, color) {
		canHold[h] = true
	}
	fmt.Println("part 1:", len(canHold))
	fmt.Println("part 2:", containedCount(bags, color))
}

func holders(bags []rules.Bag, color string) []string {
	result := make([]string, 0)
LOOP:
	for _, b := range bags {
		if b.CanHold(color) {
			result = append(result, b.Color)
			if b.CanHoldNoOthers() {
				continue LOOP
			}
			result = append(result, holders(bags, b.Color)...)
		}
	}
	return result
}

func containedCount(bags []rules.Bag, color string) (result int) {
	for _, b := range bags {
		if b.Color == color {
			for k, v := range b.Contains {
				result += v
				result += v * containedCount(bags, k)
			}
		}
	}
	return
}

