package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	pCount := n * 2

	var combinations []string
	var backtracking func(string, int, int)

	backtracking = func(prev string, op, cl int) {
		if len(prev) == pCount {
			combinations = append(combinations, prev)

			return
		}

		if len(prev)+(op-cl+2) <= pCount {
			backtracking(prev+"(", op+1, cl)
		}

		if op > cl {
			backtracking(prev+")", op, cl+1)
		}
	}

	backtracking("", 0, 0)

	return combinations
}
