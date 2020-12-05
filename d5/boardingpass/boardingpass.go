package boardingpass

import (
	"strconv"
	"strings"
)

func SeatID(code string) int {
	s := code[0:7]
	s = strings.Replace(s, "F", "0", -1)
	s = strings.Replace(s, "B", "1", -1)
	r, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		return 0
	}

	s = code[7:10]
	s = strings.Replace(s, "L", "0", -1)
	s = strings.Replace(s, "R", "1", -1)
	c, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		return 0
	}

	return int((r * 8) +  c)
}
