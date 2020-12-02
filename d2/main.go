package main

import (
	"fmt"
	"log"

	"github.com/welociraptor/adventofcode-2020/v2/d2/policy"
	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	data, err := input.Strings("d2/testdata/input")
	if err != nil {
		log.Fatal(err)
	}
	c := 0
	t := 0
	for _, i := range data {
		if valid, err := policy.Validate(i); valid && err == nil {
			c++
		}
		if valid, err := policy.ValidateToboggan(i); valid && err == nil {
			t++
		}
	}
	fmt.Printf("valid passwords old policy: %d, toboggan policy: %d", c, t)
}
