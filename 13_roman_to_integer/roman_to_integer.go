package main

func romanToInt(s string) int {
	var (
		romans = map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}
		toInt = func(r rune) int { return romans[r] }
		value = 0
	)

	for i := len(s) - 1; i >= 0; i-- {
		curRom := rune(s[i])
		curInt := toInt(curRom)

		if i != 0 && toInt(rune(s[i-1])) < curInt {
			curInt -= toInt(rune(s[i-1]))
			i--
		}

		value += curInt
	}

	return value
}
