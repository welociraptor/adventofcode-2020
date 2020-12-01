package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("d1/testdata/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]int, 0)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		data = append(data, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, i := range data {
		for _, j := range data {
			if i + j == 2020 {
				fmt.Printf("%d * %d = %d\n", i, j, i * j)
			}
			for _, k := range data {
				if i + j + k == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", i, j, k, i * j * k)
				}
			}
		}
	}
}
