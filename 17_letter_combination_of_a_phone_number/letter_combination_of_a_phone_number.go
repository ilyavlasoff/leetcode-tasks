package main

// Given a string containing digits from 2-9 inclusive, return all
// possible letter combinations that the number could represent.
// Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons)
// is given below. Note that 1 does not map to any letters.
var wordMapping = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	combinationGroups := make([][]string, len(digits))

	for i, l := range digits {
		combinationGroups[i] = wordMapping[l]
	}

	var wordCombinations []string

	var backtracking func(string, int)

	backtracking = func(res string, iter int) {
		if iter == len(digits) {
			wordCombinations = append(wordCombinations, res)

			return
		}

		currentCombination := combinationGroups[iter]
		iter++

		for _, w := range currentCombination {
			nextRes := res + w
			backtracking(nextRes, iter)
		}
	}

	backtracking("", 0)

	return wordCombinations
}
