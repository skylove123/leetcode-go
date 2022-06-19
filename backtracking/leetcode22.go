package backtracking

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	if n <= 0 {
		return result
	}
	helper(n, 0, 0, "")
	return result
}

func helper(n int, left int, right int, curr string) {
	if right == n {
		result = append(result, curr)
	}
	if left < n {
		curr += "("
		helper(n, left+1, right, curr)
		curr = curr[:len(curr)-1]
	}
	if right < left {
		curr += ")"
		helper(n, left, right+1, curr)
		curr = curr[:len(curr)-1]
	}
}
