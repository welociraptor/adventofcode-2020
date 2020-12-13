package program

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode string
	Value  int
}

func ParseProgram(input []string) []Instruction {
	program := make([]Instruction, 0)
	for _, in := range input {
		a := strings.Split(in, " ")
		v, err := strconv.Atoi(a[1])
		if err != nil {
			panic(err)
		}
		i := Instruction{
			Opcode: a[0],
			Value:  v,
		}
		program = append(program, i)
	}
	return program
}

func Run(program []Instruction) (int, bool) {
	infinite := false
	ptrack := make(map[int]bool)
	acc := 0
	for pc := 0; pc < len(program); pc++ {
		if ptrack[pc] {
			infinite = true
			break
		}
		ptrack[pc] = true
		switch program[pc].Opcode {
		case "acc":
			acc += program[pc].Value
		case "jmp":
			pc += program[pc].Value - 1
		case "nop":
		}
	}
	return acc, infinite
}