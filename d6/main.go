package main

import (
	"fmt"
	"log"

	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	data, err := input.Strings("d6/testdata/input", false)
	if err != nil {
		log.Fatal(err)
	}

	var part1 int
	var part2 int
	var groupCount int
	currentGroupAnswers := make(map[int32]int)
PARSE:
	for _, answers := range data {
		if len(answers) == 0 {
			part1 += len(currentGroupAnswers)
			for _, v := range currentGroupAnswers {
				if v == groupCount {
					part2++
				}
			}
			groupCount = 0
			currentGroupAnswers = make(map[int32]int)
			continue PARSE
		}
		groupCount++
		for _, answer := range answers {
			currentGroupAnswers[answer]++
		}
	}
	fmt.Println("part 1:", part1) // 6659
	fmt.Println("part 2:", part2) // 3219
}
