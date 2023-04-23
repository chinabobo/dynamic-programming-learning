package lecture1

/*
Problem:
	Find the sum of the first N numbers.
Objective function:
	f(i) is the sum of the first i elements.
Recurrence relation:
	f(n) = f(n-1) + n

f(0) = 0
f(1) = 0 + 1
f(2) = 0 + 1 + 2
*/

func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}
