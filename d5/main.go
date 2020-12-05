package main

import (
	"fmt"
	"log"

	"github.com/welociraptor/adventofcode-2020/v2/d5/boardingpass"
	"github.com/welociraptor/adventofcode-2020/v2/input"
)

func main() {
	data, err := input.Strings("d5/testdata/input", true)
	if err != nil {
		log.Fatal(err)
	}
	maxSeatID := 0
	seatIDs := make([]int, len(data))

	for i, bp := range data {
		sid := boardingpass.SeatID(bp)
		seatIDs[i] = sid
		if sid > maxSeatID {
			maxSeatID = sid
		}
	}

	mySeatID := 0
	for _, sid := range seatIDs {
		if !exists(seatIDs, sid + 1) && exists(seatIDs, sid + 2) {
			mySeatID = sid + 1
		}
	}
	fmt.Println(maxSeatID, mySeatID)
}

func exists(arr []int, in int) bool {
	for _, i := range arr {
		if i == in {
			return true
		}
	}
	return false
}