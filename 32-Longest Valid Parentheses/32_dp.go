package thirty_two

//FIXME Error: dp fault

func longestValidParentheses1(s string) int {

	maxLen := 0

	src := []byte(s)
	dp := make([]int, len(src))

	if len(src) < 2 {
		return 0
	}
	if src[0] == '(' && src[1] == ')' {
		dp[1] = 2
		maxLen = 2
	}

	for i := 2; i < len(src); i++ {
		if src[i] == ')' {

			if src[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else if (i-dp[i-1]) > 0 && src[i-dp[i-1]-1] == '(' {

				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-2] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-2] + 2
				}

			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}
