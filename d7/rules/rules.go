package rules

import (
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Color    string
	Contains map[string]int
}

func ParseSingle(rule string) Bag {
	colors := strings.Split(rule, " bags contain ")
	containedBagColors := strings.Split(colors[1], ", ")

	re := regexp.MustCompile(" bags?(.|,)?")

	b := Bag{
		Color:    colors[0],
		Contains: make(map[string]int),
	}

LOOP:
	for _, c := range containedBagColors {
		bag := re.ReplaceAllString(c, "")
		s := strings.Split(bag, " ")
		num, err := strconv.Atoi(s[0])
		if err != nil {
			continue LOOP
		}
		b.Contains[strings.Join(s[1:], " ")] = num
	}
	return b
}

func (b *Bag) CanHold(color string) bool {
	_, canHold := b.Contains[color]
	return canHold
}

func (b *Bag) CanHoldNoOthers() bool {
	return len(b.Contains) == 0
}
