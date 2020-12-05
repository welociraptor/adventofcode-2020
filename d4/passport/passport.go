package passport

import (
	"regexp"
	"strconv"
	"strings"
)

func Create(input []string) Passport {
	p := Passport{}
	for _, s := range input {
		ss := strings.Split(s, ":")
		switch ss[0] {
		case "byr":
			p.byr = ss[1]
		case "iyr":
			p.iyr = ss[1]
		case "eyr":
			p.eyr = ss[1]
		case "hgt":
			p.hgt = ss[1]
		case "hcl":
			p.hcl = ss[1]
		case "ecl":
			p.ecl = ss[1]
		case "pid":
			p.pid = ss[1]
		case "cid":
			p.cid = ss[1]
		}
	}
	return p
}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) Validate() bool {
	if p.byr == "" ||
		p.iyr == "" ||
		p.eyr == "" ||
		p.hgt == "" ||
		p.hcl == "" ||
		p.ecl == "" ||
		p.pid == "" {
		return false
	}
	return true
}

func (p Passport) ValidatePart2() bool {
	if !validateRange(p.byr, 1920, 2002) {
		return false
	}

	if !validateRange(p.iyr, 2010, 2020) {
		return false
	}

	if !validateRange(p.eyr, 2020, 2030) {
		return false
	}

	if !validateHeight(p.hgt) {
		return false
	}

	if !validateEyeColor(p.ecl) {
		return false
	}

	if !validateHairColor(p.hcl) {
		return false
	}

	if !validatePassportNumber(p.pid) {
		return false
	}

	return true
}

func validateRange(in string, min, max int) bool {
	i, err := strconv.Atoi(in)
	if err != nil || i < min || i > max {
		return false
	}
	return true
}

func validateHeight(hgt string) bool {
	switch {
	case strings.HasSuffix(hgt, "in"):
		hgt = strings.TrimSuffix(hgt, "in")
		return validateRange(hgt, 59, 76)
	case strings.HasSuffix(hgt, "cm"):
		hgt = strings.TrimSuffix(hgt, "cm")
		return validateRange(hgt, 150, 193)
	default:
		return false
	}
}

func validateHairColor(hcl string) bool {
	if len(hcl) != 7 {
		return false
	}
	match, err := regexp.MatchString("^#[0-9a-f]{6}$", hcl)
	if err != nil {
		return false
	}
	return match
}

var eyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validateEyeColor(ecl string) bool {
	return eyeColors[ecl]
}

func validatePassportNumber(pid string) bool {
	match, err := regexp.MatchString("^[0-9]{9}$", pid)
	if err != nil {
		return false
	}
	return match
}
