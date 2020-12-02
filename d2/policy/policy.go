package policy

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(input string) (bool, error) {
	s := strings.Split(input, " ")
	if len(s) != 3 {
		return false, fmt.Errorf("invalid input: %s", input)
	}

	min, max, err := reqCharRange(s[0])
	if err != nil {
		return false, err
	}

	c, err := reqChar(s[1])

	count := 0

	for _, i := range s[2] {
		if i == c {
			count++
		}
	}

	valid := false
	if count >= min && count <= max {
		valid = true
	}

	return valid, nil
}

func ValidateToboggan(input string) (bool, error) {
	s := strings.Split(input, " ")
	if len(s) != 3 {
		return false, fmt.Errorf("invalid input: %s", input)
	}

	p1, p2, err := reqCharRange(s[0])
	if err != nil {
		return false, err
	}

	c, err := reqChar(s[1])
	if err != nil {
		return false, err
	}

	valid := false

	if int32(s[2][p1-1]) == c && int32(s[2][p2-1]) != c {
		valid = true
	}
	if int32(s[2][p1-1]) != c && int32(s[2][p2-1]) == c {
		valid = true
	}
	return valid, nil
}

func reqCharRange(input string) (int, int, error) {
	s := strings.Split(input, "-")

	min, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid character count range: %s", s[0])
	}

	max, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid character count range: %s", s[0])
	}

	return min, max, nil
}

func reqChar(input string) (int32, error) {
	if len(input) != 2 {
		return '-', fmt.Errorf("invalid input: %s", input)
	}

	return int32(strings.TrimSuffix(input, ":")[0]), nil
}
