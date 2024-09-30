package main

import "strings"

type rti struct {
	roman   string
	integer int
}

var (
	digits = []rti{
		{roman: "I", integer: 1},
		{roman: "V", integer: 5},
		{roman: "X", integer: 10},
		{roman: "L", integer: 50},
		{roman: "C", integer: 100},
		{roman: "D", integer: 500},
		{roman: "M", integer: 1000},
	}
)

func intToRoman(num int) string {
	var (
		mult    = 1
		digit   = num % 10
		numsRev = make([]string, 10)
	)

	for num != 0 {
		numsRev = append(numsRev, getRoman(digit*mult))
		num = num / 10
		digit = num % 10
		mult *= 10
	}

	for i, j := 0, len(numsRev)-1; i < j; i, j = i+1, j-1 {
		numsRev[i], numsRev[j] = numsRev[j], numsRev[i]
	}

	return strings.Join(numsRev, "")
}

func getRoman(v int) string {
	if val, ok := getSubtractiveFormUse(v); ok {
		return val
	}

	s := ""

	for v != 0 {
		var current rti
		for i, d := range digits {
			if d.integer <= v && (i == len(digits)-1 || digits[i+1].integer > v) {
				current = d
				break
			}
		}

		s += current.roman
		v = v - current.integer
	}

	return s
}

func getSubtractiveFormUse(d int) (string, bool) {
	for k, v := range map[int]string{4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"} {
		if k == d {
			return v, true
		}
	}

	return "", false
}
