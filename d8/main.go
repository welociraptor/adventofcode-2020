package main

import (
	"fmt"
	"os"

	"github.com/welociraptor/adventofcode-2020/v2/d8/program"
	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	input, err := input.Strings("d8/testdata/input", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := program.ParseProgram(input)

	acc, _ := program.Run(p)
	fmt.Println("value of acc:", acc)

LOOP:
	for i := 0; i < len(p); i++ {
		if p[i].Opcode != "nop" && p[i].Opcode != "jmp" {
			continue LOOP
		}
		fixp := make([]program.Instruction, len(p))
		copy(fixp, p)
		switch fixp[i].Opcode {
		case "jmp":
			fixp[i].Opcode = "nop"
		case "nop":
			fixp[i].Opcode = "jmp"
		}
		acc, inf := program.Run(fixp)
		if !inf {
			fmt.Println("program fixed, acc:", acc)
			return
		}
	}
}