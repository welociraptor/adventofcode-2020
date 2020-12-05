package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/welociraptor/adventofcode-2020/v2/d4/passport"
	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	data, err := input.Strings("d4/testdata/input", false)
	if err != nil {
		log.Fatal(err)
	}

	passports := make([]passport.Passport, 0)
	ps := make([]string, 0)
PARSE:
	for _, s := range data {
		if len(s) == 0 {
			passports = append(passports, passport.Create(ps))
			ps = []string{}
			continue PARSE
		}
		ps = append(ps, strings.Split(s, " ")...)
	}
	passports = append(passports, passport.Create(ps))

	validPart1 := 0
	validPart2 := 0
	for _, pp := range passports {
		if pp.Validate() {
			validPart1++
		}
		if pp.ValidatePart2() {
			validPart2++
		}
	}

	fmt.Println("valid passports (part 1):", validPart1, "(part 2):", validPart2)
}
